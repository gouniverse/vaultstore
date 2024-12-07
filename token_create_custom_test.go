package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_TokenCreateCustom(t *testing.T) {
	db := InitDB("test_vaultstore_tokencreatecustom.db")
	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "token_create_custom",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("vault store: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()
	err = store.TokenCreateCustom(ctx, "token_custom", "test_val", "test_pass")

	if err != nil {
		t.Fatalf("vault store: Expected [err] to be nil received [%v]", err.Error())
	}

	value, err := store.TokenRead(ctx, "token_custom", "test_pass")

	if err != nil {
		t.Fatalf("vault store: Expected [err] to be nil received [%v]", err.Error())
	}

	if value != "test_val" {
		t.Fatalf("vault store: Expected [value] to be 'test_val' received [%v]", value)
	}
}
