package id

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Join metadata fields into an ID.
func Join(bucket, key string) string {
	return filepath.Join(bucket, key)
}

// Split ID into metadata fields.
func Split(id string) (string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected ID format (%q), expected: bucket/key", id)
	}

	return parts[0], parts[1], nil
}
