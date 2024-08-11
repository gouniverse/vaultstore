package vaultstore

import (
	"strings"
	"testing"
)

func Test_Store_TokenCreate(t *testing.T) {
	db := InitDB("test_vaultstore_tokencreate.db")

	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "token_create",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("value store: Expected [err] to be nil received [%v]", err.Error())
	}

	token, err := store.TokenCreate("test_val", "test_pass")

	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	if token == "" {
		t.Fatal("Token expected to not be empty")
	}

	if strings.HasPrefix(token, "tk_") == false {
		t.Fatal("Token expected to start with 'tk_' received: ", token)
	}
}
