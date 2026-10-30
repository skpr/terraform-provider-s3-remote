#!/usr/bin/make -f

export CGO_ENABLED=0

define build_step
	GOOS=$(1) GOARCH=$(2) go build -o bin/terraform-provider-s3-remote-$(1)-$(2) ./cmd/terraform-provider-s3-remote
endef

# Builds the project.
build:
	$(call build_step,linux,amd64)
	$(call build_step,linux,arm64)
	$(call build_step,darwin,amd64)
	$(call build_step,darwin,arm64)

# Run all lint checking with exit codes for CI.
lint:
	golangci-lint run ./...

# Run tests with coverage reporting.
test:
	go test -cover ./...

# Check dependencies for known vulnerabilities.
vulncheck:
	govulncheck ./...

.PHONY: build lint test vulncheck
