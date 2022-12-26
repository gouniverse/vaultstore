package vaultstore

import (
	"context"
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

// FindByID finds an entry by ID
func (st *Store) FindByID(id string) (*Vault, error) {

	sqlStr, _, errSql := goqu.Dialect(st.dbDriverName).
		From(st.vaultTableName).
		Where(goqu.Ex{"id": id}).
		Limit(1).
		ToSQL()

	if errSql != nil {
		if st.debugEnabled {
			log.Println(errSql.Error())
		}

		return nil, errSql
	}

	if st.debugEnabled {
		log.Println(sqlStr)
	}

	var vault Vault
	err := sqlscan.Get(context.Background(), st.db, &vault, sqlStr)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		}

		if sqlscan.NotFound(err) {
			return nil, nil
		}

		if st.debugEnabled {
			log.Println(err.Error())
		}

		return nil, err
	}

	return &vault, nil
}
