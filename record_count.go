package vaultstore

import (
	"context"
	"log"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	"github.com/gouniverse/base/database"
)

func (store *Store) RecordCount(ctx context.Context, options RecordQueryOptions) (int64, error) {
	options.CountOnly = true
	q := store.recordQuery(options)

	sqlStr, sqlParams, errSql := q.Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		Prepared(true).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	mapped, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
}
