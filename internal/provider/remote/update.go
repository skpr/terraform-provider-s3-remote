package remote

import (
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
		bucket = d.Get(FieldBucket).(string)
		key    = d.Get(FieldKey).(string)
		url    = d.Get(FieldURL).(string)
		hash   = d.Get(FieldHash).(string)
	)

	source, err := http.Get(url)
	if err != nil {
		return err
	}
	defer source.Body.Close()

	_, err = s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Bucket:     aws.String(bucket),
		Key:        aws.String(key),
		Body:       source.Body,
		ContentMD5: aws.String(hash),
	})
	if err != nil {
		return err
	}

	d.SetId(id.Join(bucket, key))

	return nil
}
