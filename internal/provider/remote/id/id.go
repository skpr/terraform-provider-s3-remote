package id

import (
	"fmt"
	"strings"
)

// Join metadata fields into an ID.
func Join(bucket, key string) string {
	return bucket + "/" + key
}

// Split ID into metadata fields.
func Split(id string) (string, string, error) {
	bucket, key, found := strings.Cut(id, "/")
	if !found || bucket == "" || key == "" {
		return "", "", fmt.Errorf("unexpected ID format (%q), expected: bucket/key", id)
	}

	return bucket, key, nil
}
