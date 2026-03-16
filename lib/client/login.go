package client

import (
	"context"
	"fmt"
	"strings"
)

func Login(ctx context.Context, refreshToken, clientID string) (*Client, error) {
	refreshToken = strings.TrimSpace(refreshToken)
	if refreshToken == "" {
		return nil, fmt.Errorf("refresh token is required")
	}
	clientID = strings.TrimSpace(clientID)
	if clientID == "" {
		return nil, fmt.Errorf("client ID is required")
	}

	client := New()
	client.RefreshToken = refreshToken
	client.ClientID = clientID
	if err := client.RenewToken(ctx); err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	if err := client.FetchAccountID(ctx); err != nil {
		return nil, fmt.Errorf("failed to fetch account ID: %w", err)
	}
	if err := client.SaveConfig(); err != nil {
		return nil, fmt.Errorf("failed to save config: %w", err)
	}

	return client, nil
}
