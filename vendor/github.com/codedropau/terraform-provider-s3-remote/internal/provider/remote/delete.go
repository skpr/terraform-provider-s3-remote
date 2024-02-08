package remote

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform/helper/schema"
)

// Delete the S3 object.
func Delete(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	client := s3.New(sess)

	var (
		bucket = d.Get(FieldBucket).(string)
		key    = d.Get(FieldKey).(string)
	)

	_, err := client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}
