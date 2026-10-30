package md5

import (
	"testing"
)

func TestGetHash(t *testing.T) {
	if got := GetHash([]byte("hello world")); got != "5eb63bbbe01eeed093cb22bb8f5acdc3" {
		t.Errorf("unexpected hash: %s", got)
	}
}
