package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// GetHash returns the MD5 hash of the given text.
func GetHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
