package remote

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/codedropau/terraform-provider-s3-remote/internal/provider/remote/id"
)

// Update the Project.
func Update(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	var (
		url    = d.Get(FieldBucket).(string)
		bucket = d.Get(FieldKey).(string)
		key    = d.Get(FieldURL).(string)
		hash   = d.Get(FieldHash).(string)
	)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	sum, err := md5sum(resp.Body)
	if err != nil {
		return err
	}

	if sum != hash {
		return fmt.Errorf("checksum does not match: have = %s / want = %s", sum, hash)
	}

	_, err = s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   resp.Body,
	})
	if err != nil {
		return err
	}

	d.SetId(id.Join(bucket, key))

	return nil
}

// Helper function to generate a checksum based on a stream.
func md5sum(content io.Reader) (string, error) {
	hash := md5.New()

	if _, err := io.Copy(hash, content); err != nil {
		return "", err
	}

	result := hex.EncodeToString(hash.Sum(nil))

	return result, nil
}
