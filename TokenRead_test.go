package vaultstore

import "testing"

func Test_Store_TokenRead(t *testing.T) {
	db := InitDB("test_vaultstore_tokenread.db")
	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "vault_read",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("vault store: Expected [err] to be nil received [%v]", err.Error())
	}

	id, err := store.TokenCreate("test_val", "test_pass", 20)

	if err != nil {
		t.Fatal("ValueStore Failure: ", err.Error())
	}

	val, err := store.TokenRead(id, "test_pass")
	if err != nil {
		t.Fatal("ValueRead Failure: ", err.Error())
	}

	if val != "test_val" {
		t.Fatal("ValueRetrieve Incorrect val: ", val)
	}
}
