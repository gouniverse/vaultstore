package vaultstore

import "testing"

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

	token, err := store.TokenCreate("test_val", "test_pass")
	if err != nil {
		t.Fatalf("ValueStore Failure: [%v]", err.Error())
	}

	err = store.TokenDelete(token)
	if err != nil {
		t.Fatalf("Test_Store_TokenDelete: Expected [err] to be nil received " + err.Error())
	}

	record, err := store.RecordFindByToken(token)

	if err != nil {
		t.Fatalf("Test_Store_TokenDelete: Expected [err] to be nil received [%v]", err.Error())
	}

	if record != nil {
		t.Fatalf("Test_Store_TokenDelete: Expected [record] to be nil received [%v]", record)
	}
}
