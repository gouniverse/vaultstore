package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_RecordCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("Test_Store_ValueDelete: Expected [err] to be nil received [%v]", err.Error())
	}

	record := NewRecord().SetToken("test_token1").SetValue("test_value")

	ctx := context.Background()
	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	record = NewRecord().SetToken("test_token2").SetValue("test_value")

	err = store.RecordCreate(context.Background(), *record)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	record = NewRecord().SetToken("test_token3").SetValue("test_value")

	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	count, err := store.RecordCount(ctx, RecordQueryOptions{
		Token: "test_token2",
	})

	if err != nil {
		t.Fatal("Test_Store_RecordCount: Expected [err] to be nil received " + err.Error())
	}

	if count != 1 {
		t.Fatalf("Test_Store_RecordCount: Expected [count] to be 1 received [%v]", count)
	}

}
