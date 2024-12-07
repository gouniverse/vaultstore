package vaultstore

import "context"

type StoreInterface interface {
	RecordCreate(ctx context.Context, record Record) error
	RecordFindByID(ctx context.Context, recordID string) (*Record, error)
	RecordFindByToken(ctx context.Context, token string) (*Record, error)
	RecordList(ctx context.Context, options RecordQueryOptions) ([]Record, error)
	RecordUpdate(ctx context.Context, record Record) error
	RecordDeleteByID(ctx context.Context, recordID string) error

	TokenCreate(ctx context.Context, value string, password string, tokenLength int) (token string, err error)
	TokenCreateCustom(ctx context.Context, token string, value string, password string) (err error)
	TokenDelete(ctx context.Context, token string) error
	TokenRead(ctx context.Context, token string, password string) (string, error)
	//ValueFindByID(id string) (*SearchValue, error)
	//ValueList(options SearchValueQueryOptions) ([]SearchValue, error)
	// ValueSoftDelete(valueID string) error
	// ValueSoftDeleteByID(discountID string) error
	TokenUpdate(ctx context.Context, token string, value string, password string) error
}
