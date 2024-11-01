package vaultstore

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
