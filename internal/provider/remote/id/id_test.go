package id

import (
	"testing"
)

func TestJoin(t *testing.T) {
	if got := Join("bucket", "path/to/file.zip"); got != "bucket/path/to/file.zip" {
		t.Errorf("unexpected ID: %q", got)
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		id      string
		bucket  string
		key     string
		wantErr bool
	}{
		{id: "bucket/file.zip", bucket: "bucket", key: "file.zip"},
		{id: "bucket/path/to/file.zip", bucket: "bucket", key: "path/to/file.zip"},
		{id: "bucket", wantErr: true},
		{id: "bucket/", wantErr: true},
		{id: "/file.zip", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			bucket, key, err := Split(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if bucket != tt.bucket || key != tt.key {
				t.Errorf("got (%q, %q), want (%q, %q)", bucket, key, tt.bucket, tt.key)
			}
		})
	}
}
