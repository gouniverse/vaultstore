package vaultstore

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func base64Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

func strToMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strToSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strToSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

// xorEncrypt  runs a XOR encryption on the input string
func xorEncrypt(input, key string) (output string) {
	inputBytes := []byte(input)
	keyBytes := []byte(key)
	keyLen := len(keyBytes)

	outputBytes := make([]byte, len(inputBytes))
	for i := range inputBytes {
		outputBytes[i] = inputBytes[i] ^ keyBytes[i%keyLen]
	}

	return base64Encode(outputBytes)
}

// xorDecrypt  runs a XOR encryption on the input string
func xorDecrypt(encstring string, key string) (output string, err error) {
	inputBytes, err := base64Decode(encstring)

	if err != nil {
		return "", err
	}

	outputBytes := make([]byte, len(inputBytes))
	for i, b := range inputBytes {
		outputBytes[i] = b ^ key[i%len(key)]
	}

	return string(outputBytes), nil
}
