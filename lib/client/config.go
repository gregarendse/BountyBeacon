package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".bountybeacon.json")
}

func LoadConfig() (*Client, error) {
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		return nil, err
	}
	var client Client
	if err := json.Unmarshal(data, &client); err != nil {
		return nil, err
	}
	return &client, nil
}

func LoadOrBootstrapConfig(ctx context.Context, refreshToken, clientID, apiKey string) (*Client, error) {
	client, err := LoadConfig()
	if err == nil {
		return client, nil
	}

	if refreshToken != "" && clientID != "" {
		client, err = Login(ctx, refreshToken, clientID)
		if err == nil {
			return client, nil
		}
	}

	if apiKey != "" {
		return LoginWithAPIKey(ctx, apiKey)
	}

	return nil, fmt.Errorf("config not found and no valid credentials available (set OCTOPUS_API_KEY, or OCTOPUS_REFRESH_TOKEN + OCTOPUS_CLIENT_ID)")
}

func (client *Client) SaveConfig() error {
	data, err := json.MarshalIndent(client, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath(), data, 0600)
}
