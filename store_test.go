package vaultstore

import (
	"os"
	"testing"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	os.Remove(filepath) // remove database
	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

func InitStore() *Store {
	db := InitDB("test_log_store_automigrate.db")
	return &Store{
		vaultTableName:     "test_log_store_automigrate.db",
		db:                 db,
		automigrateEnabled: false,
	}
}

func TestWithAutoMigrate(t *testing.T) {
	db := InitDB("test_log_store_automigrate.db")

	s := Store{
		vaultTableName:     "log_with_automigrate_false",
		db:                 db,
		automigrateEnabled: false,
	}

	f := WithAutoMigrate(true)
	f(&s)

	if s.automigrateEnabled != true {
		t.Fatalf("automigrateEnabled: Expected [true] received [%v]", s.automigrateEnabled)
	}

	s = Store{
		vaultTableName:     "log_with_automigrate_true",
		db:                 db,
		automigrateEnabled: true,
	}

	f = WithAutoMigrate(false)
	f(&s)

	if s.automigrateEnabled == true {
		t.Fatalf("automigrateEnabled: Expected [true] received [%v]", s.automigrateEnabled)
	}
}

func TestWithDb(t *testing.T) {
	s := Store{
		vaultTableName:     "log_with_automigrate_true",
		db:                 nil,
		automigrateEnabled: true,
	}

	db := InitDB("test")
	f := WithDb(db)
	f(&s)

	if s.db == nil {
		t.Fatalf("DB: Expected Initialized DB, received [%v]", s.db)
	}

}

func TestWithTableName(t *testing.T) {
	s := Store{
		vaultTableName:     "",
		db:                 nil,
		automigrateEnabled: false,
	}
	table_name := "Table1"
	f := WithTableName(table_name)
	f(&s)
	if s.vaultTableName != table_name {
		t.Fatalf("Expected logTableName [%v], received [%v]", table_name, s.vaultTableName)
	}
	table_name = "Table2"
	f = WithTableName(table_name)
	f(&s)
	if s.vaultTableName != table_name {
		t.Fatalf("Expected logTableName [%v], received [%v]", table_name, s.vaultTableName)
	}
}

func Test_Store_AutoMigrate(t *testing.T) {
	db := InitDB("test_log_store_automigrate.db")

	s, _ := NewStore(WithDb(db), WithTableName("log_with_automigrate"), WithAutoMigrate(true))

	s.AutoMigrate()

	if s.vaultTableName != "log_with_automigrate" {
		t.Fatalf("Expected logTableName [log_with_automigrate] received [%v]", s.vaultTableName)
	}
	if s.db == nil {
		t.Fatalf("DB Init Failure")
	}
	if s.automigrateEnabled != true {
		t.Fatalf("Failure:  WithAutoMigrate")
	}
}

func Test_Store_ValueStore(t *testing.T) {
	db := InitDB("test_log_store_automigrate.db")
	s, err := NewStore(WithDb(db), WithTableName("log_with_automigrate"), WithAutoMigrate(true))
	_, err = s.ValueStore("test_val", "test_pass")
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}
}

func Test_Store_ValueRetrieve(t *testing.T) {
	db := InitDB("test_log_store_automigrate.db")
	s, err := NewStore(WithDb(db), WithTableName("log_with_automigrate"), WithAutoMigrate(true))
	id, err := s.ValueStore("test_val", "test_pass")
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	val, err := s.ValueRetrieve(id, "test_pass")
	if err != nil {
		t.Fatalf("ValueRetrieve Failure: [%v]", err.Error())
	}

	if val != "test_val" {
		t.Fatalf("ValueRetrieve Incorrect val [%v]", val)
	}
}

func Test_Store_ValueDelete(t *testing.T) {
	db := InitDB("test_log_store_automigrate.db")
	s, err := NewStore(WithDb(db), WithTableName("log_with_automigrate"), WithAutoMigrate(true))
	id, err := s.ValueStore("test_val", "test_pass")
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	ok := s.ValueDelete(id)
	if !ok {
		t.Fatalf("ValueDelete Failure")
	}

}

func Test_decode(t *testing.T) {
	test_val := "test_value"
	test_pass := "test_password"
	encoded_str := encode(test_val, test_pass)

	str, err := decode(encoded_str, test_pass)
	if err != nil {
		t.Fatalf("decode Failure [%v]", err.Error())
	}
	if str != test_val {
		t.Fatalf("decoded String Match Failure: Expected [%v], received [%v]", test_val, str)
	}
}

func Test_encode(t *testing.T) {
	test_val := "test_value"
	test_pass := "test_password"
	encoded_str := encode(test_val, test_pass)

	str, err := decode(encoded_str, test_pass)
	if err != nil {
		t.Fatalf("encode Failure [%v]", err.Error())
	}
	if str != test_val {
		t.Fatalf("encoded String Match Failure: Expected [%v], received [%v]", test_val, str)
	}

}

func Test_createRandomBlock(t *testing.T) {
	s := createRandomBlock(10)
	if len(s) != 10 {
		t.Fatalf("createRandomBlock Error")
	}

	s = createRandomBlock(50)
	if len(s) != 50 {
		t.Fatalf("createRandomBlock Error")
	}
}

func Test_calculateRequiredBlockLength(t *testing.T) {
	i := calculateRequiredBlockLength(1000)
	if i != 1024 {
		t.Fatalf("calculateRequiredBlockLength Error")
	}
}

func Test_base64Encode(t *testing.T) {
	str := "testing"
	s := base64Encode([]byte(str))
	if len(s) == 0 {
		t.Fatalf("base64Encode Failure")
	}
}

func Test_base64Decode(t *testing.T) {
	str := "testing"
	s := base64Encode([]byte(str))
	data, err := base64Decode(s)
	if err != nil {
		t.Fatalf("base64Decode Failure: err[%v]", err.Error())
	}
	if str != string(data) {
		t.Fatalf("base64Decode Failure")
	}
}

func Test_xorEncrypt(t *testing.T) {
	str := xorEncrypt("input", "key")
	if len(str) == 0 {
		t.Fatalf("xorEncrypt Failure")
	}
}

func Test_xorDecrypt(t *testing.T) {
	str := xorEncrypt("input", "key")
	out, err := xorDecrypt(str, "key")
	if err != nil {
		t.Fatalf("xorDecrypt Failure")
	}
	if out != "input" {
		t.Fatalf("xorDecrypt Failure: Expected [input] Received [%v]", out)
	}
}

func Test_strToMD5Hash(t *testing.T) {
	ret := strToMD5Hash("testing")
	if len(ret) == 0 {
		t.Fatalf("strToMD5Hash Failure")
	}
}

func Test_strToSHA1Hash(t *testing.T) {
	ret := strToSHA1Hash("testing")
	if len(ret) == 0 {
		t.Fatalf("strToSHA1Hash Failure")
	}
}

func Test_strToSHA256Hash(t *testing.T) {
	ret := strToSHA256Hash("testing")
	if len(ret) == 0 {
		t.Fatalf("strToSHA256Hash Failure")
	}
}

func Test_isBase64(t *testing.T) {
	// Base64 of Hello -> SGVsbG8=
	ret := isBase64("SGVsbG8=")
	if !ret {
		t.Fatalf("isBase64 should ret TRUE, Failure")
	}

	ret = isBase64("Hello")
	if ret {
		t.Fatalf("isBase64 should ret FALSE, Failure")
	}
}
