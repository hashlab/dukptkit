package lib

import (
	"encoding/hex"
	"strings"
)

// GenerateCombinedKey is responsible for combining 3 component keys in order to create a combined key
func GenerateCombinedKey(key1, key2, key3 []byte) ([]byte, string) {
	xored := make([]byte, 16)
	length := len(xored)

	for i := 0; i < length; i++ {
		xored[i] = key1[i] ^ key2[i] ^ key3[i]
	}

	return xored, strings.ToUpper(hex.EncodeToString(xored))
}
