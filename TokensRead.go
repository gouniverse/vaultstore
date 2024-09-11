package vaultstore

import (
	"errors"
	"strings"

	"github.com/samber/lo"
)

// ValueRetrieve retrieves a value of a vault entry
func (st *Store) TokensRead(tokens []string, password string) (values map[string]string, err error) {
	values = map[string]string{}

	entries, err := st.RecordList(RecordQueryOptions{
		TokenIn: tokens,
	})

	if err != nil {
		return values, err
	}

	if len(entries) != len(tokens) {
		var entryTokens = lo.Map(entries, func(entry Record, _ int) string {
			return entry.Token()
		})

		_, missingTokens := lo.Difference(tokens, entryTokens)

		return values, errors.New("missing tokens: " + strings.Join(missingTokens, ", "))
	}

	for _, entry := range entries {
		decoded, err := decode(entry.Value(), password)

		if err != nil {
			return map[string]string{}, errors.New("decode error for token: " + entry.Token() + " : " + err.Error())
		}

		values[entry.Token()] = decoded
	}

	return values, nil
}
