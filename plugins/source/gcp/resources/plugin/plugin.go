// Code generated by codegen; DO NOT EDIT directly.

package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/bigquery"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/cloudbilling"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/cloudfunctions"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/compute"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/container"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/dns"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/domains"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/iam"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/logging"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/monitoring"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/redis"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/secretmanager"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/serviceusage"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/sqladmin"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/storage"
)

var (
	Version = "development"
)

const exampleConfig = `
kind: source
spec:
  tables: ["*"]
  spec:
    # Optional. List of folders to get projects from. Required permission: resourcemanager.projects.list
    # folder_ids:
    # 	- "organizations/<ORG_ID>"
    # 	- "folders/<FOLDER_ID>"
    # Optional. Maximum level of folders to recurse into
    # folders_max_depth: 5
    # Optional. If not specified either using all projects accessible.
    # project_ids:
    # 	- "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
    # Optional. ServiceAccountKeyJSON passed as value instead of a file path, can be passed also via env: CQ_SERVICE_ACCOUNT_KEY_JSON
    # service_account_key_json: <YOUR_JSON_SERVICE_ACCOUNT_KEY_DATA>
    # Optional. GRPC Retry/backoff configuration, time units in seconds. Documented in https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
    # backoff_base_delay: 1
    # backoff_multiplier: 1.6
    # backoff_max_delay: 120
    # backoff_jitter: 0.2
    # backoff_min_connect_timeout = 0
    # Optional. Max amount of retries for retrier, defaults to max 3 retries.
    # max_retries: 3
`

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"gcp",
		Version,
		[]*schema.Table{
			compute.Addresses(),
			compute.Autoscalers(),
			compute.BackendServices(),
			compute.DiskTypes(),
			compute.Disks(),
			compute.ForwardingRules(),
			compute.Instances(),
			compute.SslCertificates(),
			compute.Subnetworks(),
			compute.TargetHttpProxies(),
			compute.UrlMaps(),
			compute.VpnGateways(),
			compute.InstanceGroups(),
			compute.Images(),
			compute.Firewalls(),
			compute.Networks(),
			compute.SslPolicies(),
			compute.Interconnects(),
			compute.TargetSslProxies(),
			compute.Projects(),
			dns.Policies(),
			dns.ManagedZones(),
			domains.Registrations(),
			iam.Roles(),
			iam.ServiceAccounts(),
			container.Clusters(),
			logging.Metrics(),
			logging.Sinks(),
			redis.Instances(),
			monitoring.AlertPolicies(),
			secretmanager.Secrets(),
			serviceusage.Services(),
			sqladmin.Instances(),
			storage.Buckets(),
			cloudfunctions.Functions(),
			bigquery.Datasets(),
			cloudbilling.BillingAccounts(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
		plugins.WithClassifyError(client.ClassifyError),
	)
}
