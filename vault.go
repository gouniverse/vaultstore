package vaultstore

import (
	"time"

	"github.com/gouniverse/uid"
	"gorm.io/gorm"
)

// Vault a vault implementation
type Vault struct {
	ID        string     `gorm:"type:varchar(40);column:id;primary_key;"`
	Value     string     `gorm:"type:longtext;column:vault_value;"`
	CreatedAt time.Time  `gorm:"type:datetime;column:created_at;DEFAULT NULL;"`
	UpdatedAt time.Time  `gorm:"type:datetime;column:updated_at;DEFAULT NULL;"`
	DeletedAt *time.Time `gorm:"type:datetime;olumn:deleted_at;DEFAULT NULL;"`
}

// TableName the name of the Vault table
func (Vault) TableName() string {
	return "snv_vault"
}

// BeforeCreate adds UID to model
func (c *Vault) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uid.NanoUid()
	c.ID = uuid
	return nil
}
