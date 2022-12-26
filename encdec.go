package vaultstore

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
		return "", errors.New("base64. " + err.Error())
	}

	a := strings.Split(string(v4), "_")

	if len(a) < 2 {
		return "", errors.New("vault password incorrect")
	}

	upTo, err := strconv.Atoi(a[0])

	if err != nil {
		return "", errors.New("ATOI. " + err.Error())
	}

	v1 := a[1][0:upTo]

	v2, err := base64Decode(v1)
	if err != nil {
		return "", errors.New("Base64.2. " + err.Error())
	}

	return string(v2), nil
}

func encode(value string, password string) string {
	strongPassword := strongifyPassword(password)
	v1 := base64Encode([]byte(value))
	v2 := strconv.Itoa(len(v1)) + "_" + v1
	randomBlock := createRandomBlock(calculateRequiredBlockLength(len(v2)))
	v3 := v2 + "" + randomBlock[len(v2):len(randomBlock)]
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
	rand.Seed(time.Now().UnixNano())
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersLength := len(characters)
	randomString := ""
	for i := 0; i < length; i++ {
		randomString += string(characters[rand.Intn(charactersLength-1)])
	}
	return randomString
}

// calculateRequiredBlockLength calculates block length (128) required to contain a length
func calculateRequiredBlockLength(v int) int {
	a := 128
	for v > a {
		a = a * 2
	}
	return a
}
