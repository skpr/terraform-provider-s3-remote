package remote

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/skpr/terraform-provider-s3-remote/internal/provider/remote/id"
)

// Read the S3 object.
func Read(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	bucket, key, err := id.Split(d.Id())
	if err != nil {
		return diag.Errorf("failed to get ID: %s", err)
	}

	if err := d.Set(FieldBucket, bucket); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(FieldKey, key); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
