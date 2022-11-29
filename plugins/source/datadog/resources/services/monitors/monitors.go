// Code generated by codegen; DO NOT EDIT.

package monitors

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:      "datadog_monitors",
		Resolver:  fetchMonitors,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "deleted",
				Type:     schema.TypeTimestamp,
				Resolver: client.NullableResolver("Deleted"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: client.NullableResolver("Priority"),
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "creator",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Creator"),
			},
			{
				Name:     "message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Message"),
			},
			{
				Name:     "modified",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Modified"),
			},
			{
				Name:     "multi",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Multi"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "overall_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OverallState"),
			},
			{
				Name:     "query",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Query"),
			},
			{
				Name:     "restricted_roles",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RestrictedRoles"),
			},
			{
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			MonitorDowntimes(),
		},
	}
}
