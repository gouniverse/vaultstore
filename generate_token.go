package vaultstore

import (
	"crypto/rand"
	"fmt"
)

// GenerateToken generates a random token
// Business logic:
//  1. Generate random lowercase string
//  2. Prefix with "tk_"
func generateToken(tokenLength int) (string, error) {
	token := randomFromGamma(tokenLength-len(TOKEN_PREFIX), "abcdefghijklmnopqrstuvwxyz0123456789")

	return fmt.Sprintf("%s%s", TOKEN_PREFIX, token), nil
}

// randomFromGamma generates random string of specified length with the characters specified in the gamma string
func randomFromGamma(length int, gamma string) string {
	inRune := []rune(gamma)
	out := ""

	for i := 0; i < length+20; i++ {
		// Generate a random byte
		var b [1]byte
		if _, err := rand.Read(b[:]); err != nil {
			continue
		}

		// Map the byte to an index in the gamma string
		randomIndex := int(b[0]) % len(inRune)
		pick := inRune[randomIndex]
		out += string(pick)
	}

	return out[:length]
}
