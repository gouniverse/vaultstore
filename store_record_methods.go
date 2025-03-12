package vaultstore

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
	"github.com/samber/lo"
)

func (store *Store) RecordCount(ctx context.Context, query RecordQueryInterface) (int64, error) {
	query = query.SetCountOnly(true)
	dataset, err := query.toSelectDataset(store)

	if err != nil {
		return -1, err
	}

	sqlStr, sqlParams, errSql := dataset.Limit(1).
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

func (store *Store) RecordCreate(ctx context.Context, record RecordInterface) error {
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

// FindByID finds an entry by ID
func (st *Store) RecordFindByID(ctx context.Context, id string) (RecordInterface, error) {
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

	list := []RecordInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewRecordFromExistingData(modelMap)
		list = append(list, model)
	})

	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

// RecordFindByToken finds a record entity by token
//
// # If the supplied token is empty, an error is returned
//
// Parameters:
// - ctx: The context
// - token: The token to find
//
// Returns:
// - record: The record found
// - err: An error if something went wrong
func (st *Store) RecordFindByToken(ctx context.Context, token string) (RecordInterface, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	records, err := st.RecordList(ctx, RecordQuery().SetToken(token).SetLimit(1))

	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	return records[0], nil
}

func (store *Store) RecordUpdate(ctx context.Context, record RecordInterface) error {
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
		Where(goqu.C(COLUMN_ID).Eq(record.GetID())).
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

func (store *Store) RecordList(ctx context.Context, query RecordQueryInterface) ([]RecordInterface, error) {
	err := query.Validate()

	if err != nil {
		return []RecordInterface{}, err
	}

	dataset, err := query.toSelectDataset(store)

	if err != nil {
		return []RecordInterface{}, err
	}

	sqlStr, sqlParams, errSql := dataset.Select().Prepared(true).ToSQL()

	if errSql != nil {
		return []RecordInterface{}, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []RecordInterface{}, err
	}

	list := []RecordInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewRecordFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

// type RecordQueryOptions struct {
// 	ID          string
// 	IDIn        []string
// 	Token       string
// 	TokenIn     []string
// 	Offset      int
// 	OrderBy     string
// 	Limit       int
// 	CountOnly   bool
// 	SortOrder   string
// 	WithDeleted bool
// }

// func (store *Store) recordQuery(options RecordQueryOptions) *goqu.SelectDataset {
// 	q := goqu.Dialect(store.dbDriverName).From(store.vaultTableName)

// 	if options.ID != "" {
// 		q = q.Where(goqu.C("id").Eq(options.ID))
// 	}

// 	if options.Token != "" {
// 		q = q.Where(goqu.C(COLUMN_VAULT_TOKEN).Eq(options.Token))
// 	}

// 	if len(options.IDIn) > 0 {
// 		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn))
// 	}

// 	if len(options.TokenIn) > 0 {
// 		q = q.Where(goqu.C(COLUMN_VAULT_TOKEN).In(options.TokenIn))
// 	}

// 	if !options.CountOnly {
// 		if options.Limit > 0 {
// 			q = q.Limit(uint(options.Limit))
// 		}

// 		if options.Offset > 0 {
// 			q = q.Offset(uint(options.Offset))
// 		}
// 	}

// 	sortOrder := sb.DESC
// 	if options.SortOrder != "" {
// 		sortOrder = options.SortOrder
// 	}

// 	if options.OrderBy != "" {
// 		if strings.EqualFold(sortOrder, sb.ASC) {
// 			q = q.Order(goqu.I(options.OrderBy).Asc())
// 		} else {
// 			q = q.Order(goqu.I(options.OrderBy).Desc())
// 		}
// 	}

// 	if !options.WithDeleted {
// 		q = q.Where(goqu.C(COLUMN_DELETED_AT).Gt(carbon.Now(carbon.UTC).ToDateTimeString()))
// 	}

// 	return q
// }
