package vaultstore

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

// RecordInterface defines the methods that a Record must implement
type RecordInterface interface {
	Data() map[string]string
	DataChanged() map[string]string

	// Getters
	GetCreatedAt() string
	GetSoftDeletedAt() string
	GetID() string
	GetToken() string
	GetUpdatedAt() string
	GetValue() string

	// Setters
	SetCreatedAt(createdAt string) RecordInterface
	SetSoftDeletedAt(softDeletedAt string) RecordInterface
	SetID(id string) RecordInterface
	SetToken(token string) RecordInterface
	SetUpdatedAt(updatedAt string) RecordInterface
	SetValue(value string) RecordInterface
}

type RecordQueryInterface interface {
	Validate() error
	toSelectDataset(store StoreInterface) (*goqu.SelectDataset, error)

	IsIDSet() bool
	GetID() string
	SetID(id string) RecordQueryInterface

	IsIDInSet() bool
	GetIDIn() []string
	SetIDIn(idIn []string) RecordQueryInterface

	IsTokenSet() bool
	GetToken() string
	SetToken(token string) RecordQueryInterface

	IsTokenInSet() bool
	GetTokenIn() []string
	SetTokenIn(tokenIn []string) RecordQueryInterface

	IsOffsetSet() bool
	GetOffset() int
	SetOffset(offset int) RecordQueryInterface

	IsOrderBySet() bool
	GetOrderBy() string
	SetOrderBy(orderBy string) RecordQueryInterface

	IsLimitSet() bool
	GetLimit() int
	SetLimit(limit int) RecordQueryInterface

	IsCountOnlySet() bool
	GetCountOnly() bool
	SetCountOnly(countOnly bool) RecordQueryInterface

	IsSortOrderSet() bool
	GetSortOrder() string
	SetSortOrder(sortOrder string) RecordQueryInterface

	IsWithDeletedSet() bool
	GetWithDeleted() bool
	SetWithDeleted(withDeleted bool) RecordQueryInterface
}

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool)

	GetDbDriverName() string
	GetVaultTableName() string

	RecordCount(ctx context.Context, query RecordQueryInterface) (int64, error)
	RecordCreate(ctx context.Context, record RecordInterface) error
	RecordDeleteByID(ctx context.Context, recordID string) error
	RecordDeleteByToken(ctx context.Context, token string) error
	RecordFindByID(ctx context.Context, recordID string) (RecordInterface, error)
	RecordFindByToken(ctx context.Context, token string) (RecordInterface, error)
	RecordList(ctx context.Context, query RecordQueryInterface) ([]RecordInterface, error)
	RecordUpdate(ctx context.Context, record RecordInterface) error

	TokenCreate(ctx context.Context, value string, password string, tokenLength int) (token string, err error)
	TokenCreateCustom(ctx context.Context, token string, value string, password string) (err error)
	TokenDelete(ctx context.Context, token string) error
	TokenExists(ctx context.Context, token string) (bool, error)
	TokenRead(ctx context.Context, token string, password string) (string, error)
	TokenUpdate(ctx context.Context, token string, value string, password string) error
	TokensRead(ctx context.Context, tokens []string, password string) (map[string]string, error)
}
