package cli

import (
	"fmt"
	"strings"

	libclient "github.com/gregarendse/BountyBeacon/lib/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func loginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Log in using OCTOPUS_REFRESH_TOKEN and OCTOPUS_CLIENT_ID, then save session config",
		RunE: func(cmd *cobra.Command, args []string) error {
			token := strings.TrimSpace(viper.GetString("octopus_refresh_token"))
			if token == "" {
				return fmt.Errorf("OCTOPUS_REFRESH_TOKEN environment variable is not set\nUsage: OCTOPUS_REFRESH_TOKEN=... OCTOPUS_CLIENT_ID=... octopus-cli login")
			}
			clientID := strings.TrimSpace(viper.GetString("octopus_client_id"))
			if clientID == "" {
				return fmt.Errorf("OCTOPUS_CLIENT_ID environment variable is not set\nUsage: OCTOPUS_REFRESH_TOKEN=... OCTOPUS_CLIENT_ID=... octopus-cli login")
			}

			client, err := libclient.Login(cmd.Context(), token, clientID)
			if err != nil {
				return fmt.Errorf("login failed: %w", err)
			}

			logInfo("login successful", "account", client.AccountID)
			return nil
		},
	}
}
