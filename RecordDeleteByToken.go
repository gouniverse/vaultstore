package vaultstore

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
)

func (store *Store) RecordDeleteByToken(token string) error {
	if token == "" {
		return errors.New("token is empty")
	}

	q := goqu.Dialect(store.dbDriverName).
		Delete(store.vaultTableName).
		Prepared(true).
		Where(goqu.C(COLUMN_VAULT_TOKEN).Eq(token))

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
