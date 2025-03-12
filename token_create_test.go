package vaultstore

import (
	"context"
	"strings"
	"testing"
)

func Test_Store_TokenCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("Test_Store_TokenCreate: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()
	token, err := store.TokenCreate(ctx, "test_val", "test_pass", 20)

	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	if token == "" {
		t.Fatal("Token expected to not be empty")
	}

	if strings.HasPrefix(token, "tk_") == false {
		t.Fatal("Token expected to start with 'tk_' received: ", token)
	}

	if len(token) != 20 {
		t.Fatal("Token length expected to be 20 received: ", len(token), " token: ", token)
	}
}
