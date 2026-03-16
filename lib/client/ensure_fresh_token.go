package client

import (
	"context"
	"time"
)

const defaultTokenRefreshLeadTime = 5 * time.Minute

func (client *Client) EnsureFreshToken(ctx context.Context, refreshLeadTime time.Duration) error {
	if refreshLeadTime <= 0 {
		refreshLeadTime = defaultTokenRefreshLeadTime
	}
	if time.Until(client.ExpiresAt) > refreshLeadTime {
		return nil
	}
	return client.RenewToken(ctx)
}
