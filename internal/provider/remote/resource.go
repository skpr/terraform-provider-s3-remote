package remote

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldBucket defines the location of the file.
	FieldBucket = "bucket"
	// FieldKey defines the location of the file.
	FieldKey = "key"
	// FieldURL defines the source URL.
	FieldURL = "url"
	// FieldHash defines the source hash.
	FieldHash = "hash"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: Update,
		ReadContext:   Read,
		UpdateContext: Update,
		DeleteContext: Delete,

		Schema: map[string]*schema.Schema{
			FieldBucket: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the S3 bucket to upload the file to.",
			},
			FieldKey: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Object key to store the file under.",
			},
			FieldURL: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of the remote file to download.",
			},
			FieldHash: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Expected MD5 hash (hex encoded) of the remote file.",
			},
		},
	}
}
