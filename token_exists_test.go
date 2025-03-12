package vaultstore

import (
	"context"
	"testing"
)

func TestTokenExists(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("TestTokenExists: Expected [err] to be nil received [%v]", err.Error())
	}

	token := "token1"

	ctx := context.Background()
	exists, err := store.TokenExists(ctx, token)

	if err != nil {
		t.Fatal(err)
	}

	if exists {
		t.Fatal("token should not exist")
	}

	err = store.TokenCreateCustom(ctx, token, "value1", "password")

	if err != nil {
		t.Fatal(err)
	}

	exists, err = store.TokenExists(ctx, token)

	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("token should exist")
	}
}
