package vaultstore

import (
	"errors"
	"math/rand/v2"
	"strconv"
	"strings"
)

func decode(value string, password string) (string, error) {
	strongPassword := strongifyPassword(password)
	first, err := xorDecrypt(value, strongPassword)

	if err != nil {
		return "", errors.New("xor. " + err.Error())
	}

	if !isBase64(first) {
		return "", errors.New("vault password incorrect")
	}

	v4, err := base64Decode(first)

	if err != nil {
		return "", errors.New("base64.1. " + err.Error())
	}

	parts := strings.Split(string(v4), "_")

	if len(parts) < 2 {
		return "", errors.New("vault password incorrect")
	}

	upTo, err := strconv.Atoi(parts[0])

	if err != nil {
		return "", errors.New("atoi. " + err.Error())
	}

	after := strings.Join(parts[1:], "_")

	v1 := after[0:upTo]

	v2, err := base64Decode(v1)
	if err != nil {
		return "", errors.New("base64.2. " + err.Error())
	}

	return string(v2), nil
}

func encode(value string, password string) string {
	strongPassword := strongifyPassword(password)
	v1 := base64Encode([]byte(value))
	v2 := strconv.Itoa(len(v1)) + "_" + v1
	randomBlock := createRandomBlock(calculateRequiredBlockLength(len(v2)))
	v3 := v2 + "" + randomBlock[len(v2):]
	v4 := base64Encode([]byte(v3))
	last := xorEncrypt(v4, strongPassword)
	return last
}

// strongifyPassword Performs multiple calculations
// on top of the password and changes it to a derivative
// long hash. This is done so that even simple and not-long
// passwords  can become longer and stronger (144 characters).
func strongifyPassword(password string) string {
	p1 := strToMD5Hash(password) + strToMD5Hash(password) + strToMD5Hash(password) + strToMD5Hash(password)

	p1 = strToSHA256Hash(p1)
	p2 := strToSHA1Hash(p1) + strToSHA1Hash(p1) + strToSHA1Hash(p1)
	p3 := strToSHA1Hash(p2) + strToMD5Hash(p2) + strToSHA1Hash(p2)
	p4 := strToSHA256Hash(p3)
	p5 := strToSHA1Hash(p4) + strToSHA256Hash(strToMD5Hash(p4)) + strToSHA256Hash(strToSHA1Hash(p4)) + strToMD5Hash(p4)
	return p5
}

// createRandomBlock returns a random string of specified length
func createRandomBlock(length int) string {
	const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.IntN(len(characters))]
	}
	return string(result)
}

// calculateRequiredBlockLength calculates block length (128) required to contain a length
func calculateRequiredBlockLength(v int) int {
	a := 128
	for v > a {
		a = a * 2
	}
	return a
}
