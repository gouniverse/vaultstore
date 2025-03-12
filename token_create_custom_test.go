package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_TokenCreateCustom(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("Test_Store_TokenCreateCustom: Expected [err] to be nil received [%v]", err.Error())
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
