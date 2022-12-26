package vaultstore

import "encoding/base64"

func base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func base64Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
