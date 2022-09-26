// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityProfiles() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_security_profiles",
		Resolver:  fetchIotSecurityProfiles,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "targets",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotSecurityProfileTargets,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotSecurityProfileTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "additional_metrics_to_retain",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AdditionalMetricsToRetain"),
			},
			{
				Name:     "additional_metrics_to_retain_v2",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalMetricsToRetainV2"),
			},
			{
				Name:     "alert_targets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AlertTargets"),
			},
			{
				Name:     "behaviors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Behaviors"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "last_modified_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedDate"),
			},
			{
				Name:     "security_profile_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityProfileDescription"),
			},
			{
				Name:     "security_profile_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityProfileName"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
