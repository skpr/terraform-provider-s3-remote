package remote

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/skpr/terraform-provider-s3-remote/internal/md5"
	"github.com/skpr/terraform-provider-s3-remote/internal/provider/remote/id"
)

// Update downloads the remote file, verifies its hash and uploads it to S3.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cfg := m.(aws.Config)

	var (
		bucket   = d.Get(FieldBucket).(string)
		key      = d.Get(FieldKey).(string)
		url      = d.Get(FieldURL).(string)
		hashWant = d.Get(FieldHash).(string)
	)

	body, err := download(ctx, url)
	if err != nil {
		return diag.FromErr(err)
	}

	hashHave := md5.GetHash(body)

	if hashHave != hashWant {
		return diag.Errorf("hash mismatch: %s != %s", hashHave, hashWant)
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.Join(bucket, key))

	return nil
}

// download returns the body of the file at the given URL.
func download(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get source: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get source: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read source body: %w", err)
	}

	return body, nil
}
