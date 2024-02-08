package remote

import (
	"github.com/hashicorp/terraform/helper/schema"
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
		Create: Update,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			FieldBucket: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			FieldKey: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			FieldURL: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldHash: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
