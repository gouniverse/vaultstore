package vaultstore

import (
	"context"
	"errors"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/gouniverse/base/database"
	"github.com/samber/lo"
)

// FindByID finds an entry by ID
func (st *Store) RecordFindByID(ctx context.Context, id string) (*Record, error) {
	if id == "" {
		return nil, errors.New("record id is empty")
	}

	sqlStr, _, errSql := goqu.Dialect(st.dbDriverName).
		From(st.vaultTableName).
		Where(goqu.Ex{COLUMN_ID: id}).
		Limit(1).
		ToSQL()

	if errSql != nil {
		if st.debugEnabled {
			log.Println(errSql.Error())
		}

		return nil, errSql
	}

	if st.debugEnabled {
		log.Println(sqlStr)
	}

	modelMaps, err := database.SelectToMapString(st.toQuerableContext(ctx), sqlStr)

	if err != nil {
		return nil, err
	}

	list := []Record{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewRecordFromExistingData(modelMap)
		list = append(list, *model)
	})

	if len(list) == 0 {
		return nil, nil
	}

	return &list[0], nil
}
