package remote

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/codedropau/terraform-provider-s3-remote/internal/provider/remote/id"
)

// Read the S3 object.
func Read(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	client := s3.New(sess)

	bucket, key, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	url := d.Get(FieldBucket).(string)

	resp, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	d.Set(FieldBucket, bucket)
	d.Set(FieldKey, key)
	d.Set(FieldURL, url)
	d.Set(FieldHash, *resp.ETag)

	return nil
}
