package vaultstore

import "testing"

func Test_xorEncrypt(t *testing.T) {
	str := xorEncrypt("input", "key")
	if len(str) == 0 {
		t.Fatalf("xorEncrypt Failure")
	}
}

func Test_xorDecrypt(t *testing.T) {
	str := xorEncrypt("input", "key")
	out, err := xorDecrypt(str, "key")
	if err != nil {
		t.Fatalf("xorDecrypt Failure")
	}
	if out != "input" {
		t.Fatalf("xorDecrypt Failure: Expected [input] Received [%v]", out)
	}
}

func Test_xorEncryptDecryptLargeInput(t *testing.T) {
	var input = createRandomBlock(1000000)

	if len(input) != 1000000 {
		t.Fatalf("createRandomBlock Failure")
	}

	str := xorEncrypt(input, "key")

	out, err := xorDecrypt(str, "key")

	if err != nil {
		t.Fatalf("xorDecrypt Failure")
	}

	if out != input {
		t.Fatalf("xorDecrypt Failure: Expected [input] Received [%v]", out)
	}
}
