package vaultstore

import (
	"context"
	"testing"
)

func Test_TokensRead(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("Test_TokensRead: Expected [err] to be nil received [%v]", err.Error())
	}

	values := []string{"value1", "value2", "value3"}
	tokens := []string{"", "", ""}

	ctx := context.Background()
	for i := 0; i < len(values); i++ {
		token, err := store.TokenCreate(ctx, values[i], "test_pass", 20)

		if err != nil {
			t.Fatal("ValueStore Failure: ", err.Error())
		}

		tokens[i] = token
	}

	vals, err := store.TokensRead(ctx, tokens, "test_pass")

	if err != nil {
		t.Fatal("ValueRead Failure: ", err.Error())
	}

	for i := 0; i < len(values); i++ {
		if vals[tokens[i]] != values[i] {
			t.Fatal("ValueRetrieve Incorrect val: ", vals[tokens[i]])
		}
	}
}
