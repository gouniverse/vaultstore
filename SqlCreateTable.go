package vaultstore

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
