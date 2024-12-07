package vaultstore

import "context"

// TokenDelete deletes a token from the store
func (st *Store) TokenDelete(ctx context.Context, token string) error {
	return st.RecordDeleteByToken(ctx, token)
}
