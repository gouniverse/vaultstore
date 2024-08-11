package vaultstore

type StoreInterface interface {
	RecordCreate(record Record) error
	RecordFindByID(recordID string) (*Record, error)
	RecordFindByToken(token string) (*Record, error)
	RecordList(options RecordQueryOptions) ([]Record, error)
	RecordUpdate(Record) error
	RecordDeleteByID(recordID string) error

	TokenCreate(value string, password string, tokenLength int) (token string, err error)
	TokenCreateCustom(token string, value string, password string) (err error)
	TokenDelete(token string) error
	TokenRead(token string, password string) (string, error)
	//ValueFindByID(id string) (*SearchValue, error)
	//ValueList(options SearchValueQueryOptions) ([]SearchValue, error)
	// ValueSoftDelete(valueID string) error
	// ValueSoftDeleteByID(discountID string) error
	TokenUpdate(token string, value string, password string) error
}
