package vaultstore

import (
	"context"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/gouniverse/base/database"
)

func (store *Store) RecordDeleteByID(ctx context.Context, recordID string) error {
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

	_, err = database.Execute(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return err
	}

	return nil
}
