package lib

import (
	b "bytes"
	"crypto/des"
	"encoding/hex"
)

// HexToBytes is responsible for converting hex strings into an array of bytes
func HexToBytes(str string) ([]byte, error) {
	data, err := hex.DecodeString(str)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// IsOddParityAdjusted is responsible for checking if an array of bytes is adjusted with odd parity
func IsOddParityAdjusted(bytes []byte) bool {
	temp := make([]byte, len(bytes))

	copy(temp, bytes)

	adjustOddParity(temp)

	return b.Equal(bytes, temp)
}

func adjustOddParity(bytes []byte) {
	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		bytes[i] = (b & 0xfe) | ((((b >> 1) ^ (b >> 2) ^ (b >> 3) ^ (b >> 4) ^ (b >> 5) ^ (b >> 6) ^ (b >> 7)) ^ 0x01) & 0x01)
	}
}

func buildTdesKey(key []byte) []byte {
	var finalKey []byte

	if len(key) == 24 {
		finalKey = key
	} else if len(key) == 16 {
		finalKey = make([]byte, 24)
		copy(finalKey, key)
		copy(finalKey[16:], key[:8])
	}

	return finalKey
}

func encrypt(dst, data, key []byte) error {
	block, err := des.NewTripleDESCipher(buildTdesKey(key))

	if err != nil {
		return err
	}

	block.Encrypt(dst, data)

	return nil
}
