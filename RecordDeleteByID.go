package vaultstore

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
)

func (store *Store) RecordDeleteByID(recordID string) error {
	if recordID == "" {
		return errors.New("record id is empty")
	}

	q := goqu.Dialect(store.dbDriverName).
		Delete(store.vaultTableName).
		Prepared(true).
		Where(goqu.C(COLUMN_ID).Eq(recordID))

	sqlStr, sqlParams, err := q.ToSQL()

	if err != nil {
		return err
	}

	_, err = store.db.Exec(sqlStr, sqlParams...)

	if err != nil {
		return err
	}

	return nil
}
