package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch resources from configured providers",
	Long: `Fetch resources from configured providers

  This requires a config.hcl file which can be generated by "cloudquery init"
	`,
	Example: `  # Fetch configured providers to PostgreSQL as configured in config.hcl
  cloudquery fetch`,
	Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
		failOnError := viper.GetBool("fail-on-error")
		return c.Fetch(ctx, failOnError)
	}),
}

func init() {
	fetchCmd.SetUsageTemplate(usageTemplateWithFlags)
	fetchCmd.PersistentFlags().Bool("fail-on-error", false, "CloudQuery should return a failure error code if provider has any error")
	_ = viper.BindPFlag("fail-on-error", fetchCmd.PersistentFlags().Lookup("fail-on-error"))
	fetchCmd.Flags().Bool("skip-schema-upgrade", false, "skip schema upgrade of provider fetch, disabling this flag might cause issues")
	_ = viper.BindPFlag("skip-schema-upgrade", fetchCmd.Flags().Lookup("skip-schema-upgrade"))
	rootCmd.AddCommand(fetchCmd)
}
