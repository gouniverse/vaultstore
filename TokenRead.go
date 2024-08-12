package vaultstore

import "errors"

// ValueRetrieve retrieves a value of a vault entry
func (st *Store) TokenRead(token string, password string) (value string, err error) {
	entry, err := st.RecordFindByToken(token)

	if err != nil {
		return "", err
	}

	if entry == nil {
		return "", errors.New("value does not exist")
	}

	decoded, err := decode(entry.Value(), password)

	if err != nil {
		return "", err
	}

	return decoded, nil
}
