// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RoleServicesLastAccessed() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_role_services_last_accessed",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchIamRoleServicesLastAccessed,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "job_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver(`JobId`),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver(`ResourceARN`),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver(`ServiceName`),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceNamespace"),
			},
			{
				Name:     "last_authenticated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastAuthenticated"),
			},
			{
				Name:     "last_authenticated_entity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastAuthenticatedEntity"),
			},
			{
				Name:     "last_authenticated_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastAuthenticatedRegion"),
			},
			{
				Name:     "total_authenticated_entities",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalAuthenticatedEntities"),
			},
			{
				Name:     "tracked_actions_last_accessed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrackedActionsLastAccessed"),
			},
			{
				Name:     "entities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Entities"),
			},
		},
	}
}
