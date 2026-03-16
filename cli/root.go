package cli

import (
	"fmt"
	"strings"
	"time"

	libclient "github.com/gregarendse/BountyBeacon/lib/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func rootCommand(ctx *commandContext) *cobra.Command {
	root := &cobra.Command{
		Use:           "octopus-cli",
		Short:         "Monitor and claim Octoplus rewards",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := configureLogger(); err != nil {
				return err
			}

			if cmd.Name() == "login" {
				return nil
			}

			client, err := libclient.LoadOrBootstrapConfig(
				cmd.Context(),
				strings.TrimSpace(viper.GetString("octopus_refresh_token")),
				strings.TrimSpace(viper.GetString("octopus_client_id")),
			)
			if err != nil {
				return fmt.Errorf("unable to load session: %w\nRun: OCTOPUS_REFRESH_TOKEN=... OCTOPUS_CLIENT_ID=... octopus-cli login", err)
			}

			if err := client.EnsureFreshToken(cmd.Context(), 5*time.Minute); err != nil {
				return fmt.Errorf("session expired and refresh failed: %w", err)
			}

			ctx.client = client
			return nil
		},
	}

	root.PersistentFlags().String("log-level", "info", "Log level: debug, info, warn, error")
	root.PersistentFlags().String("log-format", "text", "Log format: text or json")
	_ = viper.BindPFlag("log_level", root.PersistentFlags().Lookup("log-level"))
	_ = viper.BindPFlag("log_format", root.PersistentFlags().Lookup("log-format"))

	root.AddCommand(loginCommand())
	root.AddCommand(rewardsCommand(ctx))
	root.AddCommand(checkCommand(ctx))
	root.AddCommand(claimCommand(ctx))
	root.AddCommand(watchCommand(ctx))
	return root
}
