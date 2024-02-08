package remote

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/skpr/terraform-provider-s3-remote/internal/md5"
	"github.com/skpr/terraform-provider-s3-remote/internal/provider/remote/id"
)

// Update the Project.
func Update(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	var (
		bucket   = d.Get(FieldBucket).(string)
		key      = d.Get(FieldKey).(string)
		url      = d.Get(FieldURL).(string)
		hashWant = d.Get(FieldHash).(string)
	)

	source, err := http.Get(url)
	if err != nil {
		return err
	}

	if source.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get source: %s", source.Status)
	}

	bodyBytes, err := io.ReadAll(source.Body)
	if err != nil {
		return fmt.Errorf("failed to read source body: %w", err)
	}

	err = source.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to close source body: %w", err)
	}

	bodyString := string(bodyBytes)

	hashHave := md5.GetHash(bodyString)

	if hashHave != hashWant {
		return fmt.Errorf("hash mismatch: %s != %s", hashHave, hashWant)
	}

	_, err = s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(bodyBytes),
	})
	if err != nil {
		return err
	}

	d.SetId(id.Join(bucket, key))

	return nil
}
