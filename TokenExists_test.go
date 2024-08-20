package vaultstore

import (
	"testing"
)

func TestTokenExists(t *testing.T) {
	db := InitDB("test_vaultstore_tokendelete.db")
	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "token_delete",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("Test_Store_ValueDelete: Expected [err] to be nil received [%v]", err.Error())
	}

	token := "token1"

	exists, err := store.TokenExists(token)

	if err != nil {
		t.Fatal(err)
	}

	if exists {
		t.Fatal("token should not exist")
	}

	err = store.TokenCreateCustom(token, "value1", "password")

	if err != nil {
		t.Fatal(err)
	}

	exists, err = store.TokenExists(token)

	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("token should exist")
	}
}
