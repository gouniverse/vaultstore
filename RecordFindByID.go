package vaultstore

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
)

// FindByID finds an entry by ID
func (st *Store) RecordFindByID(id string) (*Record, error) {

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

	db := sb.NewDatabase(st.db, st.dbDriverName)
	modelMaps, err := db.SelectToMapString(sqlStr)
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
