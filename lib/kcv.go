package lib

import (
	"encoding/hex"
	"strings"
)

// CalculateKcv is responsible for calculating the Key Checksum Value of a key
func CalculateKcv(key []byte) ([]byte, string, error) {
	kcv := make([]byte, 8)

	if err := encrypt(kcv, []byte{0, 0, 0, 0, 0, 0, 0, 0}, key); err != nil {
		return nil, "", err
	}

	return kcv[:3], strings.ToUpper(hex.EncodeToString(kcv[:3])), nil
}
