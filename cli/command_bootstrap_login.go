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
		Short: "Use saved session when valid, otherwise log in from env credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := libclient.LoadConfig()
			if err == nil {
				if err := client.EnsureFreshToken(cmd.Context(), 5*time.Minute); err == nil {
					logInfo("bootstrap login skipped", "reason", "existing session is valid", "account", client.AccountID)
					return nil
				}
				logWarn("existing session invalid, attempting env bootstrap", "error", err)
			}

			// Try refresh-token login
			token, clientID, credErr := readLoginCredentials()
			if credErr == nil {
				client, err = libclient.Login(cmd.Context(), token, clientID)
				if err == nil {
					logInfo("bootstrap login successful (refresh token)", "account", client.AccountID)
					return nil
				}
				logWarn("refresh token login failed, trying API key", "error", err)
			}

			// Fall back to API key
			apiKey, credErr := readAPIKey()
			if credErr != nil {
				return fmt.Errorf("bootstrap login failed: no valid credentials available")
			}

			client, err = libclient.LoginWithAPIKey(cmd.Context(), apiKey)
			if err != nil {
				return fmt.Errorf("bootstrap login failed: %w", err)
			}

			logInfo("bootstrap login successful (API key)", "account", client.AccountID)
			return nil
		},
	}
}
