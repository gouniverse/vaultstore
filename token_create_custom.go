package vaultstore

import (
	"context"

	"github.com/dromara/carbon/v2"
)

func (store *Store) TokenCreateCustom(ctx context.Context, token string, data string, password string) (err error) {
	encodedData := encode(data, password)

	var newEntry = NewRecord().
		SetToken(token).
		SetValue(encodedData).
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	err = store.RecordCreate(ctx, *newEntry)

	if err != nil {
		return err
	}

	return nil
}
