package randstr

import (
	"crypto/rand"
	"encoding/hex"
)

func Hex(size int) string {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
