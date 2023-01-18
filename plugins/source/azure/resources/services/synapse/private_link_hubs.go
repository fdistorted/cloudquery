package synapse

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PrivateLinkHubs() *schema.Table {
	return &schema.Table{
		Name:        "azure_synapse_private_link_hubs",
		Resolver:    fetchPrivateLinkHubs,
		Description: "https://learn.microsoft.com/en-us/rest/api/synapse/private-link-hubs/list?tabs=HTTP#privatelinkhub",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_synapse_private_link_hubs", client.Namespacemicrosoft_synapse),
		Transform:   transformers.TransformWithStruct(&armsynapse.PrivateLinkHub{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchPrivateLinkHubs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsynapse.NewPrivateLinkHubsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
