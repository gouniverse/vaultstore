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

// SqlCreateTable returns a SQL string for creating the setting table
func (st *Store) SqlCreateTable() string {
	sqlMysql := `
	CREATE TABLE IF NOT EXISTS ` + st.vaultTableName + ` (
	  id varchar(40) NOT NULL PRIMARY KEY,
	  vault_value longtext NOT NULL,
	  created_at datetime NOT NULL,
	  updated_at datetime,
	  deleted_at datetime
	);
	`

	sqlPostgres := `
	CREATE TABLE IF NOT EXISTS "` + st.vaultTableName + `" (
	  "id" varchar(40) NOT NULL PRIMARY KEY,
	  "vault_value" longtext NOT NULL,
	  "created_at" timestamptz(6) NOT NULL,
	  "updated_at" datetime,
	  "deleted_at" timestamptz(6) 
	)
	`

	sqlSqlite := `
	CREATE TABLE IF NOT EXISTS "` + st.vaultTableName + `" (
	  "id" varchar(40) NOT NULL PRIMARY KEY,
	  "vault_value" longtext NOT NULL,
	  "created_at" datetime NOT NULL,
	  "updated_at" datetime,
	  "deleted_at" datetime 
	)
	`

	sql := "unsupported driver '" + st.dbDriverName + "'"

	if st.dbDriverName == "mysql" {
		sql = sqlMysql
	}
	if st.dbDriverName == "postgres" {
		sql = sqlPostgres
	}
	if st.dbDriverName == "sqlite" {
		sql = sqlSqlite
	}

	return sql
}
