package vaultstore

import (
	"time"

	"database/sql"

	"github.com/gouniverse/uid"
)

// Vault a vault implementation
type Vault struct {
	ID        string     `json:"id" db:"id" gorm:"type:varchar(40);column:id;primary_key;"`
	Value     string     `json:"vault_value" db:"vault_value" gorm:"type:longtext;column:vault_value;"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" gorm:"type:datetime;column:created_at;DEFAULT NULL;"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" gorm:"type:datetime;column:updated_at;DEFAULT NULL;"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at" gorm:"type:datetime;olumn:deleted_at;DEFAULT NULL;"`
}

// TableName the name of the Vault table
func (Vault) TableName() string {
	return "snv_vault"
}

// BeforeCreate adds UID to model
func (c *Vault) BeforeCreate(tx *sql.DB) (err error) {
	uuid := uid.NanoUid()
	c.ID = uuid
	return nil
}
