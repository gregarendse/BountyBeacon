package cli

import (
	"fmt"
	"time"

	libclient "github.com/gregarendse/BountyBeacon/lib/client"
	"github.com/spf13/cobra"
)

func bootstrapLoginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "bootstrap-login",
		Short: "Use saved session when valid, otherwise log in from OCTOPUS_REFRESH_TOKEN and OCTOPUS_CLIENT_ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := libclient.LoadConfig()
			if err == nil {
				if err := client.EnsureFreshToken(cmd.Context(), 5*time.Minute); err == nil {
					logInfo("bootstrap login skipped", "reason", "existing session is valid", "account", client.AccountID)
					return nil
				}
				logWarn("existing session invalid, attempting env bootstrap", "error", err)
			}

			token, clientID, err := readLoginCredentials()
			if err != nil {
				return err
			}

			client, err = libclient.Login(cmd.Context(), token, clientID)
			if err != nil {
				return fmt.Errorf("bootstrap login failed: %w", err)
			}

			logInfo("bootstrap login successful", "account", client.AccountID)
			return nil
		},
	}
}
