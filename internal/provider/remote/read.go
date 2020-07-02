package remote

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/codedropau/terraform-provider-s3-remote/internal/provider/remote/id"
)

// Read the S3 object.
func Read(d *schema.ResourceData, m interface{}) error {
	bucket, key, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	var (
		url  = d.Get(FieldBucket).(string)
		hash = d.Get(FieldHash).(string)
	)

	d.Set(FieldBucket, bucket)
	d.Set(FieldKey, key)
	d.Set(FieldURL, url)
	d.Set(FieldHash, hash)

	return nil
}
