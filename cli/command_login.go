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
		Short: "Log in using credentials (API key or refresh token), then save session config",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Try refresh-token login first
			token, clientID, err := readLoginCredentials()
			if err == nil {
				client, loginErr := libclient.Login(cmd.Context(), token, clientID)
				if loginErr == nil {
					logInfo("login successful (refresh token)", "account", client.AccountID)
					return nil
				}
				logWarn("refresh token login failed, trying API key", "error", loginErr)
			}

			// Fall back to API key
			apiKey, credErr := readAPIKey()
			if credErr != nil {
				if err != nil {
					return fmt.Errorf("login failed: no valid credentials available\nSet OCTOPUS_API_KEY, or OCTOPUS_REFRESH_TOKEN + OCTOPUS_CLIENT_ID")
				}
				return credErr
			}

			client, loginErr := libclient.LoginWithAPIKey(cmd.Context(), apiKey)
			if loginErr != nil {
				return fmt.Errorf("login failed: %w", loginErr)
			}

			logInfo("login successful (API key)", "account", client.AccountID)
			return nil
		},
	}
}

func readLoginCredentials() (string, string, error) {
	token := strings.TrimSpace(viper.GetString("octopus_refresh_token"))
	if token == "" {
		return "", "", fmt.Errorf("OCTOPUS_REFRESH_TOKEN is not set")
	}
	clientID := strings.TrimSpace(viper.GetString("octopus_client_id"))
	if clientID == "" {
		return "", "", fmt.Errorf("OCTOPUS_CLIENT_ID is not set")
	}

	return token, clientID, nil
}

func readAPIKey() (string, error) {
	apiKey := strings.TrimSpace(viper.GetString("octopus_api_key"))
	if apiKey == "" {
		return "", fmt.Errorf("OCTOPUS_API_KEY is not set")
	}
	return apiKey, nil
}

