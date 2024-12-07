package vaultstore

import (
	"context"
	"errors"
)

// RecordFindByToken finds a store record by token
func (st *Store) RecordFindByToken(ctx context.Context, token string) (*Record, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	records, err := st.RecordList(ctx, RecordQueryOptions{
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	return &records[0], nil
}
