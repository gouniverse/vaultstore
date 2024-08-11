package vaultstore

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"strings"
)

// GenerateToken generates a random token
// Business logic:
//  1. Generate random lowercase string
//  2. Prefix with "tk_"
func GenerateToken() (string, error) {
	tokenLength := 26 // Adjust token length as needed (without prefix)
	b := make([]byte, tokenLength)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	encoded := strings.ToLower(base32.StdEncoding.EncodeToString(b))
	return fmt.Sprintf("tk_%s", encoded), nil
}
