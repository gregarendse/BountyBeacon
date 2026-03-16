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

func LoadOrBootstrapConfig(ctx context.Context, refreshToken, clientID string) (*Client, error) {
	client, err := LoadConfig()
	if err == nil {
		return client, nil
	}

	if refreshToken == "" {
		return nil, fmt.Errorf("config not found and OCTOPUS_REFRESH_TOKEN is not set")
	}
	if clientID == "" {
		return nil, fmt.Errorf("config not found and OCTOPUS_CLIENT_ID is not set")
	}

	return Login(ctx, refreshToken, clientID)
}

func (client *Client) SaveConfig() error {
	data, err := json.MarshalIndent(client, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath(), data, 0600)
}
