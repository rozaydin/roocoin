package util

import (
	"crypto"

	_ "golang.org/x/crypto/sha3"
)

// ...
func CalculateSHA3_256Hash(input string) []byte {
	hashFunc := crypto.SHA3_256.New()
	hashFunc.Write([]byte(input))
	calculatedHash := hashFunc.Sum(nil)
	return calculatedHash
}
