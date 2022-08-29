// Code generated by codegen; DO NOT EDIT.

package domains

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Registrations() *schema.Table {
	return &schema.Table{
		Name:      "gcp_domains_registrations",
		Resolver:  fetchRegistrations,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "contact_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ContactSettings"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "dns_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DnsSettings"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpireTime"),
			},
			{
				Name:     "issues",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Issues"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "management_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ManagementSettings"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "pending_contact_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingContactSettings"),
			},
			{
				Name:     "register_failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegisterFailureReason"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "supported_privacy",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedPrivacy"),
			},
			{
				Name:     "transfer_failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransferFailureReason"),
			},
		},
	}
}

func fetchRegistrations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Domains.Projects.Locations.Registrations.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Registrations

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
