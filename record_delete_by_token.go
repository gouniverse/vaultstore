package vaultstore

import (
	"context"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/gouniverse/base/database"
)

func (store *Store) RecordDeleteByToken(ctx context.Context, token string) error {
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

	_, err = database.Execute(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return err
	}

	return nil
}
