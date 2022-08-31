// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func AddOns() *schema.Table {
	return &schema.Table{
		Name:        "heroku_add_ons",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#add-on-attributes",
		Resolver:    fetchAddOns,
		Columns: []schema.Column{
			{
				Name:     "actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Actions"),
			},
			{
				Name:     "addon_service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddonService"),
			},
			{
				Name:     "app",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("App"),
			},
			{
				Name:     "billed_price",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BilledPrice"),
			},
			{
				Name:     "billing_entity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BillingEntity"),
			},
			{
				Name:     "config_vars",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ConfigVars"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Plan"),
			},
			{
				Name:     "provider_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProviderID"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "web_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebURL"),
			},
		},
	}
}

func fetchAddOns(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AddOnList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
