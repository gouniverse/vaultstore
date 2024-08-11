package vaultstore

// TokenDelete deletes a token from the store
func (st *Store) TokenDelete(token string) error {
	return st.RecordDeleteByToken(token)
}
