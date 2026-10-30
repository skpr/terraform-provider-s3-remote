package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

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
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					FieldProfile: {
						Type:        schema.TypeString,
						Optional:    true,
						DefaultFunc: schema.EnvDefaultFunc("AWS_PROFILE", ""),
						Description: "AWS profile for authentication.",
					},
					FieldRegion: {
						Type:        schema.TypeString,
						Optional:    true,
						DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", ""),
						Description: "AWS region where the bucket is located.",
					},
				},
				ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
					var opts []func(*config.LoadOptions) error

					if v, ok := d.GetOk(FieldRegion); ok {
						opts = append(opts, config.WithRegion(v.(string)))
					}

					if v, ok := d.GetOk(FieldProfile); ok {
						opts = append(opts, config.WithSharedConfigProfile(v.(string)))
					}

					cfg, err := config.LoadDefaultConfig(ctx, opts...)
					if err != nil {
						return nil, diag.FromErr(err)
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
