// Code generated by codegen; DO NOT EDIT.

package servicequotas

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Quotas() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicequotas_quotas",
		Description: "https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceQuota.html",
		Resolver:    fetchServicequotasQuotas,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QuotaArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "adjustable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Adjustable"),
			},
			{
				Name:     "error_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ErrorReason"),
			},
			{
				Name:     "global_quota",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("GlobalQuota"),
			},
			{
				Name:     "period",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Period"),
			},
			{
				Name:     "quota_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QuotaCode"),
			},
			{
				Name:     "quota_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QuotaName"),
			},
			{
				Name:     "service_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceCode"),
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
			{
				Name:     "usage_metric",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UsageMetric"),
			},
			{
				Name:     "value",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Value"),
			},
		},
	}
}
