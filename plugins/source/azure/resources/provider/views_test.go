package provider

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/views"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestViews(t *testing.T) {
	providertest.HelperTestView(t, providertest.ViewTestCase{
		Provider: Provider(),
		SQLView:  views.ResourceView,
	})
}
