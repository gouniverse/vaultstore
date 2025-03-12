package vaultstore

import (
	"errors"
	"os"
	"testing"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) (*sql.DB, error) {
	if filepath != ":memory:" && fileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			return nil, err
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initStore(filepath string) (StoreInterface, error) {
	db, err := initDB(filepath)
	if err != nil {
		return nil, err
	}

	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "vault_token",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("unexpected nil store")
	}

	return store, nil
}

func TestWithAutoMigrateFalse(t *testing.T) {
	db, err := initDB(":memory:")

	if err != nil {
		t.Fatalf("initDB: Expected [err] to be nil received [%v]", err.Error())
	}

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

func Test_Store_AutoMigrate(t *testing.T) {
	db, err := initDB(":memory:")

	if err != nil {
		t.Fatalf("initDB: Expected [err] to be nil received [%v]", err.Error())
	}

	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "vault_automigrate",
		DB:                 db,
		AutomigrateEnabled: false,
	})

	if err != nil {
		t.Fatalf("automigrateEnabled: Expected [err] to be nill received [%v]", err.Error())
	}

	if store.automigrateEnabled != false {
		t.Fatalf("automigrateEnabled: Expected [false] received [%v]", store.automigrateEnabled)
	}

	err = store.AutoMigrate()

	if err != nil {
		t.Fatalf("AutoMigrate Failure [%v]", err.Error())
	}

	if store.vaultTableName != "vault_automigrate" {
		t.Fatalf("Expected vaultTableName [vault_automigrate] received [%v]", store.vaultTableName)
	}
	if store.db == nil {
		t.Fatalf("DB Init Failure")
	}
	if store.automigrateEnabled != false {
		t.Fatalf("Failure:  AutoMigrate")
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
