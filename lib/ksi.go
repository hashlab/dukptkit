package lib

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// GenerateKSI is responsible for generating a BDK identifier (KSI) with a given number of bytes
func GenerateKSI(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(bytes)), nil
}
