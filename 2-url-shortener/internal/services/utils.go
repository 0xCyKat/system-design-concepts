package services

import (
	"crypto/sha256"
	"encoding/hex"
)

func getHash(longUrl string) string {
	data := []byte(longUrl)
	hashBytes := sha256.Sum256(data)
	hash := hex.EncodeToString(hashBytes[:])

	return hash[:7]
}
