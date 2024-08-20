package vaultstore

import "errors"

func (store *Store) TokenExists(token string) (bool, error) {
	if token == "" {
		return false, errors.New("token is empty")
	}

	count, err := store.RecordCount(RecordQueryOptions{
		Token: token,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
