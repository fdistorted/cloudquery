// Code generated by codegen; DO NOT EDIT.

package domains

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"

	"github.com/cloudquery/plugin-sdk/schema"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_domains",
		Resolver: fetchDomains,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "ttl",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TTL"),
			},
			{
				Name:     "zone_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ZoneFile"),
			},
		},

		Relations: []*schema.Table{
			Records(),
		},
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {

	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	for !done {
		listFunc := func() error {
			data, resp, err := svc.Services.Domains.List(ctx, opt)
			if err != nil {
				return errors.WithStack(err)
			}
			// pass the current page's data to our result channel
			res <- data
			// if we are at the last page, break out the for loop
			if resp.Links == nil || resp.Links.IsLastPage() {
				done = true
				return nil
			}
			page, err := resp.Links.CurrentPage()
			if err != nil {
				return errors.WithStack(err)
			}
			// set the page we want for the next request
			opt.Page = page + 1
			return nil
		}

		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
