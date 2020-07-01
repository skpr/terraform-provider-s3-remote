Terraform Provider: S3 Remote
=============================

Terraform provider for deploying a remote file to AWS S3

## Usage

```hcl
resource "aws_s3_bucket_object_remote" "function" {
  bucket = "lambda-functions"
  key    = "example.zip"

  url  = "https://github.com/codedropau/example/releases/download/v0.0.1/function.zip"
  hash = "xxxxxxxxxxxxxxxxxxxxxxxx"
}
```
