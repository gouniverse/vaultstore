package vaultstore

import "context"

// RecordInterface defines the methods that a Record must implement
type RecordInterface interface {
	Data() map[string]string
	DataChanged() map[string]string

	// Getters
	GetCreatedAt() string
	GetDeletedAt() string
	GetID() string
	GetToken() string
	GetUpdatedAt() string
	GetValue() string

	// Setters
	SetCreatedAt(createdAt string) RecordInterface
	SetDeletedAt(deletedAt string) RecordInterface
	SetID(id string) RecordInterface
	SetToken(token string) RecordInterface
	SetUpdatedAt(updatedAt string) RecordInterface
	SetValue(value string) RecordInterface
}

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool)

	RecordCount(ctx context.Context, options RecordQueryOptions) (int64, error)
	RecordCreate(ctx context.Context, record RecordInterface) error
	RecordDeleteByID(ctx context.Context, recordID string) error
	RecordDeleteByToken(ctx context.Context, token string) error
	RecordFindByID(ctx context.Context, recordID string) (RecordInterface, error)
	RecordFindByToken(ctx context.Context, token string) (RecordInterface, error)
	RecordList(ctx context.Context, options RecordQueryOptions) ([]RecordInterface, error)
	RecordUpdate(ctx context.Context, record RecordInterface) error

	TokenCreate(ctx context.Context, value string, password string, tokenLength int) (token string, err error)
	TokenCreateCustom(ctx context.Context, token string, value string, password string) (err error)
	TokenDelete(ctx context.Context, token string) error
	TokenExists(ctx context.Context, token string) (bool, error)
	TokenRead(ctx context.Context, token string, password string) (string, error)
	TokenUpdate(ctx context.Context, token string, value string, password string) error
	TokensRead(ctx context.Context, tokens []string, password string) (map[string]string, error)
}
