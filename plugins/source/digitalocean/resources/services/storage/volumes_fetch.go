package storage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"
)

func fetchVolumes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	opt := &godo.ListVolumeParams{
		ListOptions: &godo.ListOptions{PerPage: client.MaxItemsPerPage},
	}

	done := false
	listFunc := func() error {
		data, resp, err := svc.Services.Storage.ListVolumes(ctx, opt)
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
		opt.ListOptions.Page = page + 1
		return nil
	}

	for !done {
		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
