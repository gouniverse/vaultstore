package vaultstore

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
)

func (store *Store) RecordUpdate(record Record) error {
	record.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	dataChanged := record.DataChanged()

	delete(dataChanged, "id")   // ID is not updateable
	delete(dataChanged, "hash") // Hash is not updateable

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

	_, err = store.db.Exec(sqlStr, sqlParams...)
	if err != nil {
		return err
	}

	return nil
}
