package vaultstore

import (
	"crypto/sha1"
	"encoding/hex"
)

func strToSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
