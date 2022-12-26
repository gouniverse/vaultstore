package vaultstore

import "errors"

// ValueRetrieve retrieves a value of a vault entry
func (st *Store) ValueRetrieve(id string, password string) (value string, err error) {
	entry, errFind := st.FindByID(id)

	if errFind != nil {
		return "", err
	}

	if entry == nil {
		return "", errors.New("value does not exist")
	}

	decoded, err := decode(entry.Value, password)

	if err != nil {
		return "", err
	}

	return decoded, nil
}
