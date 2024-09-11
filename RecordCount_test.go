package vaultstore

import "testing"

func Test_Store_RecordCount(t *testing.T) {
	db := InitDB("test_vaultstore_recordcount.db")
	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "record_count",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("Test_Store_ValueDelete: Expected [err] to be nil received [%v]", err.Error())
	}

	record := NewRecord().SetToken("test_token1").SetValue("test_value")

	err = store.RecordCreate(*record)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	record = NewRecord().SetToken("test_token2").SetValue("test_value")

	err = store.RecordCreate(*record)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	record = NewRecord().SetToken("test_token3").SetValue("test_value")

	err = store.RecordCreate(*record)
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	count, err := store.RecordCount(RecordQueryOptions{
		Token: "test_token2",
	})

	if err != nil {
		t.Fatalf("Test_Store_RecordCount: Expected [err] to be nil received " + err.Error())
	}

	if count != 1 {
		t.Fatalf("Test_Store_RecordCount: Expected [count] to be 1 received [%v]", count)
	}

}