package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:            "cloudflare",
		Version:         Version,
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
			"access_groups":     services.AccessGroups(),
			"accounts":          services.Accounts(),
			"certificate_packs": services.CertificatePacks(),
			"dns_records":       services.DNSRecords(),
			"images":            services.Images(),
			"ips":               services.Ips(),
			"wafs":              services.Wafs(),
			"waf_overrides":     services.WafOverrides(),
			"workers_scripts":   services.WorkersScripts(),
			"workers_routes":    services.WorkersRoutes(),
			"zones":             services.Zones(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
