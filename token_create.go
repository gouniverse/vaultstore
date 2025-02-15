package vaultstore

import (
	"context"

	"github.com/dromara/carbon/v2"
)

// TokenCreate creates a new record and returns the token
func (st *Store) TokenCreate(ctx context.Context, data string, password string, tokenLength int) (token string, err error) {
	token, err = generateToken(tokenLength)

	if err != nil {
		return "", err
	}

	encodedData := encode(data, password)

	var newEntry = NewRecord().
		SetToken(token).
		SetValue(encodedData).
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	err = st.RecordCreate(ctx, *newEntry)

	if err != nil {
		return "", err
	}

	return token, nil
}
