// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func Invoices() *schema.Table {
	return &schema.Table{
		Name:        "heroku_invoices",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#invoice-attributes",
		Resolver:    fetchInvoices,
		Columns: []schema.Column{
			{
				Name:     "charges_total",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("ChargesTotal"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "credits_total",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("CreditsTotal"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Number"),
			},
			{
				Name:     "period_end",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeriodEnd"),
			},
			{
				Name:     "period_start",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeriodStart"),
			},
			{
				Name:     "state",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "total",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Total"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchInvoices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.InvoiceList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
