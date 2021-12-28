package vaultstore

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"database/sql"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlserver"
	"github.com/georgysavva/scany/sqlscan"
)

// Store defines a session store
type Store struct {
	vaultTableName     string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debug              bool
}

// StoreOption options for the vault store
type StoreOption func(*Store)

// WithAutoMigrate sets the table name for the cache store
func WithAutoMigrate(automigrateEnabled bool) StoreOption {
	return func(s *Store) {
		s.automigrateEnabled = automigrateEnabled
	}
}

// WithGormDb sets the GORM database for the cache store
func WithDb(db *sql.DB) StoreOption {
	return func(s *Store) {
		s.db = db
		s.dbDriverName = s.DriverName(s.db)
	}
}

// WithDebug Enables debug logging
func WithDebug(debug bool) StoreOption {
	return func(s *Store) {
		s.debug = debug
	}
}

// WithTableName sets the table name for the cache store
func WithTableName(vaultTableName string) StoreOption {
	return func(s *Store) {
		s.vaultTableName = vaultTableName
	}
}

// NewStore creates a new entity store
func NewStore(opts ...StoreOption) (*Store, error) {
	store := &Store{}
	for _, opt := range opts {
		opt(store)
	}

	if store.db == nil {
		return nil, errors.New("log store: db is required")
	}

	if store.dbDriverName == "" {
		return nil, errors.New("log store: dbDriverName is required")
	}

	if store.vaultTableName == "" {
		return nil, errors.New("log store: vaultTableName is required")
	}

	if store.automigrateEnabled == true {
		store.AutoMigrate()
	}

	return store, nil
}

// AutoMigrate auto migrate
func (st *Store) AutoMigrate() {
	sql := st.SqlCreateTable()

	if st.debug {
		log.Println(sql)
	}

	_, err := st.db.Exec(sql)

	if err != nil {
		log.Println(err)
		return
	}

	return
}

// DriverName finds the driver name from database
func (st *Store) DriverName(db *sql.DB) string {
	dv := reflect.ValueOf(db.Driver())
	driverFullName := dv.Type().String()
	if strings.Contains(driverFullName, "mysql") {
		return "mysql"
	}
	if strings.Contains(driverFullName, "postgres") || strings.Contains(driverFullName, "pq") {
		return "postgres"
	}
	if strings.Contains(driverFullName, "sqlite") {
		return "sqlite"
	}
	if strings.Contains(driverFullName, "mssql") {
		return "mssql"
	}
	return driverFullName
}

// EnableDebug - enables the debug option
func (st *Store) EnableDebug(debug bool) {
	st.debug = debug
}

// FindByID finds a user by ID
func (st *Store) FindByID(id string) *Vault {

	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).From(st.vaultTableName).Where(goqu.Ex{"id": id}).Limit(1).ToSQL()

	if st.debug {
		log.Println(sqlStr)
	}

	var vault Vault
	err := sqlscan.Get(context.Background(), st.db, &vault, sqlStr)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		log.Fatal("Failed to execute query: ", err)
		return nil
	}

	return &vault

	// found, err := ds.ScanStruct(vault)

	// if err != nil {
	// 	if st.debug {
	// 		log.Println(err.Error())
	// 	}
	// 	return nil
	// }
	// if !found {
	// 	return nil
	// }

	// return vault
}

// ValueDelete removes all keys from the sessiom
func (st *Store) ValueDelete(id string) error {
	sqlStr, _, _ := goqu.Dialect(st.dbDriverName).From(st.vaultTableName).Where(goqu.C("id").Eq(id), goqu.C("deleted_at").IsNull()).Delete().ToSQL()

	if st.debug {
		log.Println(sqlStr)
	}

	_, err := st.db.Exec(sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Fatal("Failed to execute query: ", err)
		return nil
	}

	return nil
}

// ValueRetrieve retrieves a value of a vault entry
func (st *Store) ValueRetrieve(id string, password string) (value string, err error) {
	entry := st.FindByID(id)

	if entry == nil {
		return "", errors.New("Value does not exist")
	}

	decoded, err := decode(entry.Value, password)

	if err != nil {
		return "", err
	}

	return decoded, nil
}

// ValueStore creates a new vault entry and returns the ID
func (st *Store) ValueStore(value string, password string) (id string, err error) {
	encoded := encode(value, password)
	var newEntry = Vault{
		ID:        fmt.Sprintf("%v", time.Now().UnixNano()),
		Value:     encoded,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).Insert(st.vaultTableName).Rows(newEntry).ToSQL()
	if st.debug {
		log.Println(sqlStr)
	}

	_, err = st.db.Exec(sqlStr)
	if err != nil {
		if st.debug {
			log.Println(err.Error())
		}
		return "", err
	}
	return newEntry.ID, nil
}

func decode(value string, password string) (string, error) {
	strongPassword := strongifyPassword(password)
	first, err := xorDecrypt(value, strongPassword)

	if err != nil {
		return "", errors.New("XOR. " + err.Error())
	}

	if isBase64(first) == false {
		return "", errors.New("Vault password incorrect")
	}

	v4, err := base64Decode(first)

	if err != nil {
		return "", errors.New("Base64. " + err.Error())
	}

	a := strings.Split(string(v4), "_")

	if len(a) < 2 {
		return "", errors.New("Vault password incorrect")
	}

	upTo, err := strconv.Atoi(a[0])

	if err != nil {
		return "", errors.New("ATOI. " + err.Error())
	}

	v1 := a[1][0:upTo]

	v2, err := base64Decode(v1)
	if err != nil {
		return "", errors.New("Base64.2. " + err.Error())
	}

	return string(v2), nil
}

func encode(value string, password string) string {
	strongPassword := strongifyPassword(password)
	v1 := base64Encode([]byte(value))
	v2 := strconv.Itoa(len(v1)) + "_" + v1
	randomBlock := createRandomBlock(calculateRequiredBlockLength(len(v2)))
	v3 := v2 + "" + randomBlock[len(v2):len(randomBlock)]
	v4 := base64Encode([]byte(v3))
	last := xorEncrypt(v4, strongPassword)
	return last
}

// strongifyPassword Performs multiple calculations on top of the password and changes it to a derivative long hash. This is done so that even simple and not-long passwords  can become longer and stronger (144 characters).
func strongifyPassword(password string) string {
	p1 := strToMD5Hash(password) + strToMD5Hash(password) + strToMD5Hash(password) + strToMD5Hash(password)

	p1 = strToSHA256Hash(p1)
	p2 := strToSHA1Hash(p1) + strToSHA1Hash(p1) + strToSHA1Hash(p1)
	p3 := strToSHA1Hash(p2) + strToMD5Hash(p2) + strToSHA1Hash(p2)
	p4 := strToSHA256Hash(p3)
	p5 := strToSHA1Hash(p4) + strToSHA256Hash(strToMD5Hash(p4)) + strToSHA256Hash(strToSHA1Hash(p4)) + strToMD5Hash(p4)
	return p5
}

// createRandomBlock returns a random string of specified length
func createRandomBlock(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersLength := len(characters)
	randomString := ""
	for i := 0; i < length; i++ {
		randomString += string(characters[rand.Intn(charactersLength-1)])
	}
	return randomString
}

// calculateRequiredBlockLength calculates block length (128) required to contain a length
func calculateRequiredBlockLength(v int) int {
	a := 128
	for v > a {
		a = a * 2
	}
	return a
}

func base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func base64Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

// xorEncrypt  runs a XOR encryption on the input string
func xorEncrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return base64Encode([]byte(output))
}

// xorDecrypt  runs a XOR encryption on the input string
func xorDecrypt(encstring string, key string) (output string, err error) {
	inputBytes, err := base64Decode(encstring)

	if err != nil {
		return "", err
	}

	input := string(inputBytes)

	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output, nil
}

func strToMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strToSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strToSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func isBase64(value string) bool {
	base64Regex := "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	rxBase64 := regexp.MustCompile(base64Regex)
	return rxBase64.MatchString(value)
}
