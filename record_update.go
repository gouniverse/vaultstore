package vaultstore

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
)

func (store *Store) RecordUpdate(ctx context.Context, record Record) error {
	record.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	dataChanged := record.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable
	delete(dataChanged, "hash")    // Hash is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, sqlParams, err := goqu.Dialect(store.dbDriverName).
		Update(store.vaultTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(record.ID())).
		ToSQL()

	if err != nil {
		return err
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	_, err = database.Execute(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return err
	}

	return nil
}
