Terraform Provider: S3 Remote
=============================

Terraform provider which downloads a remote file, verifies its MD5 hash and uploads it to an AWS S3 bucket.

This is useful for deploying release artifacts (e.g. Lambda function packages published as GitHub release assets) into S3
without having to download them as part of your pipeline.

## Installation

This provider is not published to the Terraform Registry. Download the binary for your platform from the
[releases page](https://github.com/skpr/terraform-provider-s3-remote/releases) and place it in a
[filesystem mirror](https://developer.hashicorp.com/terraform/cli/config/config-file#filesystem_mirror), e.g.

```
~/.terraform.d/plugins/registry.terraform.io/skpr/s3-remote/<VERSION>/<OS>_<ARCH>/terraform-provider-s3-remote_v<VERSION>
```

Then declare it in your configuration:

```hcl
terraform {
  required_providers {
    s3remote = {
      source  = "skpr/s3-remote"
      version = "<VERSION>"
    }
  }
}
```

## Usage

```hcl
provider "s3remote" {
  region = "ap-southeast-2"
}

resource "aws_s3_bucket_object_remote" "function" {
  # Required because the resource type prefix ("aws") does not match this provider's name.
  provider = s3remote

  bucket = "lambda-functions"
  key    = "example.zip"

  url  = "https://github.com/codedropau/example/releases/download/v0.0.1/function.zip"
  hash = "xxxxxxxxxxxxxxxxxxxxxxxx"
}
```

The `hash` is the hex encoded MD5 checksum of the remote file, which can be generated with:

```bash
curl -sL <URL> | md5sum
```

The file is re-uploaded whenever `url` or `hash` changes. Changing `bucket` or `key` replaces the object.

### Provider Arguments

| Argument  | Description                                    | Default               |
|-----------|------------------------------------------------|-----------------------|
| `region`  | AWS region where the bucket is located.        | `AWS_REGION` env var  |
| `profile` | AWS shared config profile for authentication.  | `AWS_PROFILE` env var |

Credentials are resolved using the standard
[AWS SDK credential chain](https://docs.aws.amazon.com/sdkref/latest/guide/standardized-credentials.html)
(environment variables, shared config, SSO, instance/task roles etc.).

### Resource Arguments: `aws_s3_bucket_object_remote`

| Argument | Description                                         | Required |
|----------|-----------------------------------------------------|----------|
| `bucket` | Name of the S3 bucket to upload the file to.        | Yes      |
| `key`    | Object key to store the file under.                 | Yes      |
| `url`    | URL of the remote file to download.                 | Yes      |
| `hash`   | Expected MD5 hash (hex encoded) of the remote file. | Yes      |

## Development

Requires Go (see `go.mod` for the minimum version).

```bash
make build      # Build binaries into ./bin
make test       # Run tests
make lint       # Run golangci-lint
make vulncheck  # Check dependencies for known vulnerabilities (govulncheck)
```

Releases are published by GoReleaser when a `v*` tag is pushed.
