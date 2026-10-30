package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// GetHash returns the MD5 hash of the given data.
func GetHash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
