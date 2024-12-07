package vaultstore

import "strings"

func IsToken(s string) bool {
	return strings.HasPrefix(s, TOKEN_PREFIX)
}
