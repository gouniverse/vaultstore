package vaultstore

import (
	"errors"
)

// TokenUpdate updates a token in the store
func (st *Store) TokenUpdate(token string, value string, password string) (err error) {
	entry, errFind := st.RecordFindByToken(token)

	if errFind != nil {
		return err
	}

	if entry == nil {
		return errors.New("token does not exist")
	}

	encodedValue := encode(value, password)

	entry.SetValue(encodedValue)

	return st.RecordUpdate(*entry)
}
