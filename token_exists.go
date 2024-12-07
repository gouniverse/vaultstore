package vaultstore

import (
	"context"
	"errors"
)

func (store *Store) TokenExists(ctx context.Context, token string) (bool, error) {
	if token == "" {
		return false, errors.New("token is empty")
	}

	count, err := store.RecordCount(ctx, RecordQueryOptions{
		Token: token,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
