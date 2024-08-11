package vaultstore

// RecordFindByToken finds a store record by token
func (st *Store) RecordFindByToken(token string) (*Record, error) {
	records, err := st.RecordList(RecordQueryOptions{
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
