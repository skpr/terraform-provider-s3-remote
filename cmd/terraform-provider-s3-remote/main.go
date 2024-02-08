package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	"github.com/skpr/terraform-provider-s3-remote/internal/provider/remote"
)

const (
	// ResourceBucketObjectRemote provides a resource for storing remote files in S3.
	ResourceBucketObjectRemote = "aws_s3_bucket_object_remote"

	// FieldRegion identifier for region field.
	FieldRegion = "region"
	// FieldProfile identifier for profile field.
	FieldProfile = "profile"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					FieldProfile: {
						Type:        schema.TypeString,
						Optional:    true,
						DefaultFunc: schema.EnvDefaultFunc("AWS_PROFILE", ""),
						Description: "AWS Profile for authentication.",
					},
					FieldRegion: {
						Type:        schema.TypeString,
						Optional:    true,
						DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", ""),
						Description: "AWS Profile for authentication.",
					},
				},
				ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
					cfg := &aws.Config{}

					if v, ok := d.GetOk(FieldRegion); ok {
						cfg.Region = aws.String(v.(string))
					}

					if v, ok := d.GetOk(FieldProfile); ok {
						cfg.Credentials = credentials.NewSharedCredentials("", v.(string))
					}

					return session.NewSession(cfg)
				},
				ResourcesMap: map[string]*schema.Resource{
					ResourceBucketObjectRemote: remote.Resource(),
				},
			}
		},
	})
}
