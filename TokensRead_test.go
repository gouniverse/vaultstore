package vaultstore

import (
	"testing"
)

func Test_TokensRead(t *testing.T) {
	db := InitDB("test_vaultstore_tokenread.db")
	store, err := NewStore(NewStoreOptions{
		VaultTableName:     "vault_read",
		DB:                 db,
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatalf("vault store: Expected [err] to be nil received [%v]", err.Error())
	}

	values := []string{"value1", "value2", "value3"}
	tokens := []string{"", "", ""}

	for i := 0; i < len(values); i++ {
		token, err := store.TokenCreate(values[i], "test_pass", 20)

		if err != nil {
			t.Fatal("ValueStore Failure: ", err.Error())
		}

		tokens[i] = token
	}

	vals, err := store.TokensRead(tokens, "test_pass")

	if err != nil {
		t.Fatal("ValueRead Failure: ", err.Error())
	}

	for i := 0; i < len(values); i++ {
		if vals[tokens[i]] != values[i] {
			t.Fatal("ValueRetrieve Incorrect val: ", vals[tokens[i]])
		}
	}
}
