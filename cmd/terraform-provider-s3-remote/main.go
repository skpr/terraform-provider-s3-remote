package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
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
					var opts []func(*config.LoadOptions) error

					if v, ok := d.GetOk(FieldRegion); ok {
						opts = append(opts, config.WithRegion(v.(string)))
					}

					if v, ok := d.GetOk(FieldProfile); ok {
						opts = append(opts, config.WithSharedConfigProfile(v.(string)))
					}

					cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
					if err != nil {
						return nil, err
					}

					return cfg, nil
				},
				ResourcesMap: map[string]*schema.Resource{
					ResourceBucketObjectRemote: remote.Resource(),
				},
			}
		},
	})
}
