package vaultstore

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
)

func (store *Store) RecordCreate(ctx context.Context, record Record) error {
	record.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	record.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := record.Data()

	sqlStr, sqlParams, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.vaultTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		if store.debugEnabled {
			log.Println(errSql.Error())
		}

		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		if store.debugEnabled {
			log.Println(err)
		}

		return err
	}

	return nil
}
