package vaultstore

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

// == CLASS ==================================================================

type Record struct {
	dataobject.DataObject
}

// == CONSTRUCTORS ===========================================================

func NewRecord() *Record {
	d := (&Record{}).
		SetID(uid.HumanUid()).
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetDeletedAt(sb.MAX_DATETIME)

	return d
}

func NewRecordFromExistingData(data map[string]string) *Record {
	o := &Record{}
	o.Hydrate(data)
	return o
}

// == METHODS ================================================================

// == SETTERS AND GETTERS ====================================================

func (v *Record) CreatedAt() string {
	return v.Get(COLUMN_CREATED_AT)
}

func (v *Record) SetCreatedAt(createdAt string) *Record {
	v.Set(COLUMN_CREATED_AT, createdAt)
	return v
}

func (v *Record) DeletedAt() string {
	return v.Get(COLUMN_DELETED_AT)
}

func (v *Record) SetDeletedAt(deletedAt string) *Record {
	v.Set(COLUMN_DELETED_AT, deletedAt)
	return v
}

func (v *Record) ID() string {
	return v.Get(COLUMN_ID)
}

func (v *Record) SetID(id string) *Record {
	v.Set(COLUMN_ID, id)
	return v
}

func (v *Record) Token() string {
	return v.Get(COLUMN_VAULT_TOKEN)
}

func (v *Record) SetToken(token string) *Record {
	v.Set(COLUMN_VAULT_TOKEN, token)
	return v
}

func (v *Record) UpdatedAt() string {
	return v.Get(COLUMN_UPDATED_AT)
}

func (v *Record) SetUpdatedAt(updatedAt string) *Record {
	v.Set(COLUMN_UPDATED_AT, updatedAt)
	return v
}

func (v *Record) Value() string {
	return v.Get(COLUMN_VAULT_VALUE)
}

func (v *Record) SetValue(value string) *Record {
	v.Set(COLUMN_VAULT_VALUE, value)
	return v
}

// Record a vault implementation
// type Record struct {
// 	ID        string     `json:"id" db:"id" gorm:"type:varchar(40);column:id;primary_key;"`
// 	Record     string     `json:"vault_value" db:"vault_value" gorm:"type:longtext;column:vault_value;"`
// 	CreatedAt time.Time  `json:"created_at" db:"created_at" gorm:"type:datetime;column:created_at;DEFAULT NULL;"`
// 	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" gorm:"type:datetime;column:updated_at;DEFAULT NULL;"`
// 	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at" gorm:"type:datetime;olumn:deleted_at;DEFAULT NULL;"`
// }

// TableName the name of the Record table
// func (Record) TableName() string {
// 	return "snv_vault"
// }

// // BeforeCreate adds UID to model
// func (c *Record) BeforeCreate(tx *sql.DB) (err error) {
// 	uuid := uid.NanoUid()
// 	c.ID = uuid
// 	return nil
// }
