package vaultstore

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
)

// ValueDelete removes all keys from the sessiom
func (st *Store) ValueDelete(id string) error {
	sqlStr, _, _ := goqu.Dialect(st.dbDriverName).
		From(st.vaultTableName).
		Where(goqu.C("id").Eq(id), goqu.C("deleted_at").IsNull()).
		Delete().
		ToSQL()

	if st.debugEnabled {
		log.Println(sqlStr)
	}

	_, err := st.db.Exec(sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Fatal("Failed to execute query: ", err)
		return nil
	}

	return nil
}
