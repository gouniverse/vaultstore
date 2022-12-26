package vaultstore

import (
	"errors"
)

// NewStore creates a new entity store
func NewStore(opts NewStoreOptions) (*Store, error) {
	store := &Store{
		vaultTableName:     opts.VaultTableName,
		automigrateEnabled: opts.AutomigrateEnabled,
		db:                 opts.DB,
		dbDriverName:       opts.DbDriverName,
		debugEnabled:       opts.DebugEnabled,
	}

	if store.vaultTableName == "" {
		return nil, errors.New("vault store: vaultTableName is required")
	}

	if store.db == nil {
		return nil, errors.New("vault store: DB is required")
	}

	if store.dbDriverName == "" {
		store.dbDriverName = driverName(store.db)
	}

	if store.automigrateEnabled {
		store.AutoMigrate()
	}

	return store, nil
}
