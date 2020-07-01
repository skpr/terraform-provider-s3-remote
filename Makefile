#!/usr/bin/make -f

export CGO_ENABLED=0
export GO111MODULE=on

default: lint test build

# Build for multiple OS types.
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/terraform-provider-s3-remote_{{.OS}}_{{.Arch}}' -ldflags=${LGFLAGS} github.com/codedropau/terraform-provider-s3-remote/cmd/terraform-provider-s3-remote

# Run all lint checking with exit codes for CI.
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run go fmt against code
fmt:
	go fmt ./...

# Run tests with coverage reporting.
test:
	go test -cover ./...
