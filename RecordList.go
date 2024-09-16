package vaultstore

import (
	"log"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
)

type RecordQueryOptions struct {
	ID          string
	IDIn        []string
	Token       string
	TokenIn     []string
	Offset      int
	OrderBy     string
	Limit       int
	CountOnly   bool
	SortOrder   string
	WithDeleted bool
}

func (store *Store) RecordList(options RecordQueryOptions) ([]Record, error) {
	q := store.recordQuery(options)

	sqlStr, _, errSql := q.Select().ToSQL()

	if errSql != nil {
		return []Record{}, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)
	modelMaps, err := db.SelectToMapString(sqlStr)
	if err != nil {
		return []Record{}, err
	}

	list := []Record{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewRecordFromExistingData(modelMap)
		list = append(list, *model)
	})

	return list, nil
}

func (store *Store) recordQuery(options RecordQueryOptions) *goqu.SelectDataset {
	q := goqu.Dialect(store.dbDriverName).From(store.vaultTableName)

	if options.ID != "" {
		q = q.Where(goqu.C("id").Eq(options.ID))
	}

	if options.Token != "" {
		q = q.Where(goqu.C(COLUMN_VAULT_TOKEN).Eq(options.Token))
	}

	if len(options.IDIn) > 0 {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn))
	}

	if len(options.TokenIn) > 0 {
		q = q.Where(goqu.C(COLUMN_VAULT_TOKEN).In(options.TokenIn))
	}

	if !options.CountOnly {
		if options.Limit > 0 {
			q = q.Limit(uint(options.Limit))
		}

		if options.Offset > 0 {
			q = q.Offset(uint(options.Offset))
		}
	}

	sortOrder := sb.DESC
	if options.SortOrder != "" {
		sortOrder = options.SortOrder
	}

	if options.OrderBy != "" {
		if strings.EqualFold(sortOrder, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy).Desc())
		}
	}

	if !options.WithDeleted {
		q = q.Where(goqu.C(COLUMN_DELETED_AT).Gt(carbon.Now(carbon.UTC).ToDateTimeString()))
	}

	return q
}
