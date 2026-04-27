package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func GenerateRandomString(s int) string {
	b := GenerateRandomBytes(int(math.Ceil(float64(s) / 2)))
	if s%2 == 0 {
		return hex.EncodeToString(b)
	}
	return hex.EncodeToString(b)[0:s]
}
