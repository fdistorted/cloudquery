// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Networks() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_networks",
		Resolver:  fetchNetworks,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "i_pv_4_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPv4Range"),
			},
			{
				Name:     "auto_create_subnetworks",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoCreateSubnetworks"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enable_ula_internal_ipv_6",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableUlaInternalIpv6"),
			},
			{
				Name:     "firewall_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FirewallPolicy"),
			},
			{
				Name:     "gateway_i_pv_4",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GatewayIPv4"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "internal_ipv_6_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InternalIpv6Range"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "mtu",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Mtu"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network_firewall_policy_enforcement_order",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkFirewallPolicyEnforcementOrder"),
			},
			{
				Name:     "peerings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Peerings"),
			},
			{
				Name:     "routing_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RoutingConfig"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "self_link_with_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLinkWithId"),
			},
			{
				Name:     "subnetworks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Subnetworks"),
			},
		},
	}
}

func fetchNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Networks.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
