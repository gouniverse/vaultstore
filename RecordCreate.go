package vaultstore

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-module/carbon/v2"
)

func (store *Store) RecordCreate(record Record) error {
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

	_, err := store.db.Exec(sqlStr, sqlParams...)

	if err != nil {
		if store.debugEnabled {
			log.Println(err)
		}

		return err
	}

	return nil
}
