// Code generated by codegen; DO NOT EDIT.

package quicksight

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataSources() *schema.Table {
	return &schema.Table{
		Name:      "aws_quicksight_data_sources",
		Resolver:  fetchQuicksightDataSources,
		Multiplex: client.ServiceAccountRegionMultiplexer("quicksight"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "data_source_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataSourceId"),
			},
			{
				Name:     "error_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ErrorInfo"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "secret_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecretArn"),
			},
			{
				Name:     "ssl_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SslProperties"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "vpc_connection_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConnectionProperties"),
			},
		},
	}
}
