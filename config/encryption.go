package config

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(input string) string {
	data := []byte(input)

	hashBytes := sha256.Sum256(data)
	return hex.EncodeToString(hashBytes[:])
}
