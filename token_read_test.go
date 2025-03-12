package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_TokenRead(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("Test_Store_TokenRead: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()
	id, err := store.TokenCreate(ctx, "test_val", "test_pass", 20)

	if err != nil {
		t.Fatal("ValueStore Failure: ", err.Error())
	}

	val, err := store.TokenRead(ctx, id, "test_pass")
	if err != nil {
		t.Fatal("ValueRead Failure: ", err.Error())
	}

	if val != "test_val" {
		t.Fatal("ValueRetrieve Incorrect val: ", val)
	}
}
