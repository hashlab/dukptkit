package lib

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// GenerateComponentKey is responsible for generating component keys with a given number of bytes and also with or without odd parity
func GenerateComponentKey(length int, forceOdd bool) ([]byte, string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return nil, "", err
	}

	if forceOdd {
		adjustOddParity(bytes)
	}

	return bytes, strings.ToUpper(hex.EncodeToString(bytes)), nil
}
