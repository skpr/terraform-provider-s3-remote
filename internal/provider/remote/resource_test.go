package remote

import (
	"testing"
)

func TestResource(t *testing.T) {
	if err := Resource().InternalValidate(nil, true); err != nil {
		t.Fatalf("resource schema is invalid: %s", err)
	}
}
