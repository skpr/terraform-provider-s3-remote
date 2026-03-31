package remote

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hashicorp/terraform/helper/schema"
)

// Delete the S3 object.
func Delete(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)

	client := s3.NewFromConfig(cfg)

	var (
		bucket = d.Get(FieldBucket).(string)
		key    = d.Get(FieldKey).(string)
	)

	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}
