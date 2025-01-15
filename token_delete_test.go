package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_TokenDelete(t *testing.T) {
	db := InitDB("test_vaultstore_tokendelete.db")
	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "token_delete",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("Test_Store_ValueDelete: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()
	token, err := store.TokenCreate(ctx, "test_val", "test_pass", 20)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	err = store.TokenDelete(ctx, token)
	if err != nil {
		t.Fatal("Test_Store_TokenDelete: Expected [err] to be nil received " + err.Error())
	}

	record, err := store.RecordFindByToken(ctx, token)

	if err != nil {
		t.Fatalf("Test_Store_TokenDelete: Expected [err] to be nil received [%v]", err.Error())
	}

	if record != nil {
		t.Fatalf("Test_Store_TokenDelete: Expected [record] to be nil received [%v]", record)
	}
}
