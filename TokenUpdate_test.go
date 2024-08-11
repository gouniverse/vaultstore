package vaultstore

import "testing"

func Test_Store_TokenUpdate(t *testing.T) {
	db := InitDB("test_vaultstore_tokenupdate.db")

	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "token_update",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("vault store: Expected [err] to be nil received [%v]", err.Error())
	}

	token, err := store.TokenCreate("test_val", "test_pass")

	if err != nil {
		t.Fatal("TokenCreate Failure: ", err.Error())
	}

	val, err := store.TokenRead(token, "test_pass")
	if err != nil {
		t.Fatal("TokenRead Failure: ", err.Error())
	}

	if val != "test_val" {
		t.Fatal("TokenRead Incorrect val: ", val)
	}

	err = store.TokenUpdate(token, "test_val2", "test_pass")

	if err != nil {
		t.Fatal("TokenUpdate Failure: ", err.Error())
	}

	val, err = store.TokenRead(token, "test_pass")

	if err != nil {
		t.Fatal("TokenRead Failure: ", err.Error())
	}

	if val != "test_val2" {
		t.Fatal("TokenRead Incorrect val: ", val)
	}
}
