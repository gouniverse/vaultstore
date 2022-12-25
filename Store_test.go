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

func TestWithAutoMigrate(t *testing.T) {
	db := InitDB("test_vaultstore_with_automigrate.db")

	storeAutomigrateFalse, errAutomigrateFalse := NewStore(NewStoreOptions{
		VaultTableName:     "vault_with_automigrate_false",
		DB:                 db,
		AutomigrateEnabled: false,
	})

	if errAutomigrateFalse != nil {
		t.Fatalf("automigrateEnabled: Expected [err] to be nill received [%v]", errAutomigrateFalse.Error())
	}

	if storeAutomigrateFalse.automigrateEnabled != false {
		t.Fatalf("automigrateEnabled: Expected [false] received [%v]", storeAutomigrateFalse.automigrateEnabled)
	}

	storeAutomigrateTrue, errAutomigrateTrue := NewStore(NewStoreOptions{
		VaultTableName:     "vault_with_automigrate_true",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if errAutomigrateTrue != nil {
		t.Fatalf("automigrateEnabled: Expected [err] to be nill received [%v]", errAutomigrateTrue.Error())
	}

	if storeAutomigrateTrue.automigrateEnabled != true {
		t.Fatalf("automigrateEnabled: Expected [true] received [%v]", storeAutomigrateTrue.automigrateEnabled)
	}
}

// func TestWithDb(t *testing.T) {
// 	s := Store{
// 		vaultTableName:     "vault",
// 		db:                 nil,
// 		automigrateEnabled: true,
// 	}

// 	db := InitDB("test_vaultstore_with_db")
// 	f := WithDb(db)
// 	f(&s)

// 	if s.db == nil {
// 		t.Fatalf("DB: Expected Initialized DB, received [%v]", s.db)
// 	}

// }

// func TestWithTableName(t *testing.T) {
// 	s := Store{
// 		vaultTableName:     "",
// 		db:                 nil,
// 		automigrateEnabled: false,
// 	}
// 	table_name := "Table1"
// 	f := WithTableName(table_name)
// 	f(&s)
// 	if s.vaultTableName != table_name {
// 		t.Fatalf("Expected vaultTableName [%v], received [%v]", table_name, s.vaultTableName)
// 	}
// 	table_name = "Table2"
// 	f = WithTableName(table_name)
// 	f(&s)
// 	if s.vaultTableName != table_name {
// 		t.Fatalf("Expected vaultTableName [%v], received [%v]", table_name, s.vaultTableName)
// 	}
// }

// func Test_Store_AutoMigrate(t *testing.T) {
// 	db := InitDB("test_vaultstore_automigrate.db")

// 	s, _ := NewStore(WithDb(db), WithTableName("vault_automigrate"), WithAutoMigrate(true))

// 	s.AutoMigrate()

// 	if s.vaultTableName != "vault_automigrate" {
// 		t.Fatalf("Expected vaultTableName [vault_automigrate] received [%v]", s.vaultTableName)
// 	}
// 	if s.db == nil {
// 		t.Fatalf("DB Init Failure")
// 	}
// 	if s.automigrateEnabled != true {
// 		t.Fatalf("Failure:  AutoMigrate")
// 	}
// }

// func Test_Store_ValueStore(t *testing.T) {
// 	db := InitDB("test_vaultstore_valuestore.db")
// 	s, err := NewStore(WithDb(db), WithTableName("vault_store"), WithAutoMigrate(true))
// 	_, err = s.ValueStore("test_val", "test_pass")
// 	if err != nil {
// 		t.Fatalf("ValueStore Failure: [%v]", err.Error())
// 	}
// }

// func Test_Store_ValueRetrieve(t *testing.T) {
// 	db := InitDB("test_vaultstore_valueretrieve.db")
// 	s, err := NewStore(WithDb(db), WithTableName("vault_retrieve"), WithAutoMigrate(true))
// 	id, err := s.ValueStore("test_val", "test_pass")
// 	if err != nil {
// 		t.Fatalf("ValueStore Failure: [%v]", err.Error())
// 	}

// 	val, err := s.ValueRetrieve(id, "test_pass")
// 	if err != nil {
// 		t.Fatalf("ValueRetrieve Failure: [%v]", err.Error())
// 	}

// 	if val != "test_val" {
// 		t.Fatalf("ValueRetrieve Incorrect val [%v]", val)
// 	}
// }

// func Test_Store_ValueDelete(t *testing.T) {
// 	db := InitDB("test_vaultstore_valuedelete.db")
// 	s, err := NewStore(WithDb(db), WithTableName("vault_delete"), WithAutoMigrate(true))

// 	id, err := s.ValueStore("test_val", "test_pass")
// 	if err != nil {
// 		t.Fatalf("ValueStore Failure: [%v]", err.Error())
// 	}

// 	errDelete := s.ValueDelete(id)
// 	if errDelete != nil {
// 		t.Fatalf("ValueDelete Failed: " + errDelete.Error())
// 	}

// }

// func Test_decode(t *testing.T) {
// 	test_val := "test_value"
// 	test_pass := "test_password"
// 	encoded_str := encode(test_val, test_pass)

// 	str, err := decode(encoded_str, test_pass)
// 	if err != nil {
// 		t.Fatalf("decode Failure [%v]", err.Error())
// 	}
// 	if str != test_val {
// 		t.Fatalf("decoded String Match Failure: Expected [%v], received [%v]", test_val, str)
// 	}
// }

// func Test_encode(t *testing.T) {
// 	test_val := "test_value"
// 	test_pass := "test_password"
// 	encoded_str := encode(test_val, test_pass)

// 	str, err := decode(encoded_str, test_pass)
// 	if err != nil {
// 		t.Fatalf("encode Failure [%v]", err.Error())
// 	}
// 	if str != test_val {
// 		t.Fatalf("encoded String Match Failure: Expected [%v], received [%v]", test_val, str)
// 	}

// }

// func Test_createRandomBlock(t *testing.T) {
// 	s := createRandomBlock(10)
// 	if len(s) != 10 {
// 		t.Fatalf("createRandomBlock Error")
// 	}

// 	s = createRandomBlock(50)
// 	if len(s) != 50 {
// 		t.Fatalf("createRandomBlock Error")
// 	}
// }

// func Test_calculateRequiredBlockLength(t *testing.T) {
// 	i := calculateRequiredBlockLength(1000)
// 	if i != 1024 {
// 		t.Fatalf("calculateRequiredBlockLength Error")
// 	}
// }

// func Test_base64Encode(t *testing.T) {
// 	str := "testing"
// 	s := base64Encode([]byte(str))
// 	if len(s) == 0 {
// 		t.Fatalf("base64Encode Failure")
// 	}
// }

// func Test_base64Decode(t *testing.T) {
// 	str := "testing"
// 	s := base64Encode([]byte(str))
// 	data, err := base64Decode(s)
// 	if err != nil {
// 		t.Fatalf("base64Decode Failure: err[%v]", err.Error())
// 	}
// 	if str != string(data) {
// 		t.Fatalf("base64Decode Failure")
// 	}
// }

// func Test_xorEncrypt(t *testing.T) {
// 	str := xorEncrypt("input", "key")
// 	if len(str) == 0 {
// 		t.Fatalf("xorEncrypt Failure")
// 	}
// }

// func Test_xorDecrypt(t *testing.T) {
// 	str := xorEncrypt("input", "key")
// 	out, err := xorDecrypt(str, "key")
// 	if err != nil {
// 		t.Fatalf("xorDecrypt Failure")
// 	}
// 	if out != "input" {
// 		t.Fatalf("xorDecrypt Failure: Expected [input] Received [%v]", out)
// 	}
// }

// func Test_strToMD5Hash(t *testing.T) {
// 	ret := strToMD5Hash("testing")
// 	if len(ret) == 0 {
// 		t.Fatalf("strToMD5Hash Failure")
// 	}
// }

// func Test_strToSHA1Hash(t *testing.T) {
// 	ret := strToSHA1Hash("testing")
// 	if len(ret) == 0 {
// 		t.Fatalf("strToSHA1Hash Failure")
// 	}
// }

// func Test_strToSHA256Hash(t *testing.T) {
// 	ret := strToSHA256Hash("testing")
// 	if len(ret) == 0 {
// 		t.Fatalf("strToSHA256Hash Failure")
// 	}
// }

// func Test_isBase64(t *testing.T) {
// 	// Base64 of Hello -> SGVsbG8=
// 	ret := isBase64("SGVsbG8=")
// 	if !ret {
// 		t.Fatalf("isBase64 should ret TRUE, Failure")
// 	}

// 	ret = isBase64("Hello")
// 	if ret {
// 		t.Fatalf("isBase64 should ret FALSE, Failure")
// 	}
// }
