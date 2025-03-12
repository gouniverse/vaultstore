package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_RecordCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("Test_Store_RecordCreate: Expected [err] to be nil received: ", err.Error())
	}

	record := NewRecord().SetToken("test_token").SetValue("test_value")

	ctx := context.Background()
	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatal("Test_Store_RecordCreate: Expected [err] to be nil received " + err.Error())
	}

	exists, err := store.TokenExists(ctx, "test_token")

	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("Test_Store_RecordCreate: token should exist")
	}
}
