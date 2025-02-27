// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func Peerings() *schema.Table {
	return &schema.Table{
		Name:        "heroku_peerings",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#peering-attributes",
		Resolver:    fetchPeerings,
		Columns: []schema.Column{
			{
				Name:     "aws_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsAccountID"),
			},
			{
				Name:     "aws_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsRegion"),
			},
			{
				Name:     "aws_vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsVpcID"),
			},
			{
				Name:     "cidr_blocks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CIDRBlocks"),
			},
			{
				Name:     "expires",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Expires"),
			},
			{
				Name:     "pcx_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PcxID"),
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
		},
	}
}

func fetchPeerings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.Space, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.SpaceList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		items = append(items, v...)
	}

	for _, it := range items {
		nextRange = &heroku.ListRange{
			Field: "id",
			Max:   1000,
		}
		// Roundtripper middleware in client/pagination.go
		// sets the nextRange value after each request
		for nextRange.Max != 0 {
			ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
			v, err := c.Heroku.PeeringList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
