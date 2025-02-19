package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/cloudquery/plugins/source/gcp/resources/plugin"
)

const sentryDsn = "https://c30e57a331fe4101a11b3c83d780793f@o1396617.ingest.sentry.io/6720365"

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
		SentryDsn:    sentryDsn,
	})
}
