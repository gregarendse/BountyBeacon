package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gregarendse/BountyBeacon/lib/operations"
)

func LoginWithAPIKey(ctx context.Context, apiKey string) (*Client, error) {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	payload := operations.NewObtainKrakenTokenRequest(apiKey)
	resp, err := MakeUnauthenticatedRequest[ObtainKrakenTokenVariables, ObtainKrakenTokenData](ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("obtainKrakenToken request failed: %w", err)
	}

	if len(resp.Errors) > 0 {
		return nil, fmt.Errorf("obtainKrakenToken returned errors: %s", resp.Errors[0].Message)
	}

	result := resp.Data.ObtainKrakenToken
	if result.Token == "" {
		return nil, fmt.Errorf("obtainKrakenToken returned empty token")
	}

	expiresAt, err := parseJWTExpiry(result.Token)
	if err != nil {
		// Fall back to a conservative 1-hour expiry if we can't parse the JWT
		expiresAt = time.Now().Add(1 * time.Hour)
	}

	client := New()
	client.AccessToken = result.Token
	client.RefreshToken = result.RefreshToken
	client.ExpiresAt = expiresAt

	if err := client.FetchAccountID(ctx); err != nil {
		return nil, fmt.Errorf("failed to fetch account ID: %w", err)
	}
	if err := client.SaveConfig(); err != nil {
		return nil, fmt.Errorf("failed to save config: %w", err)
	}

	return client, nil
}

// parseJWTExpiry extracts the "exp" claim from a JWT without verifying the signature.
func parseJWTExpiry(token string) (time.Time, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid JWT: expected 3 parts, got %d", len(parts))
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to decode JWT payload: %w", err)
	}

	var claims struct {
		Exp int64 `json:"exp"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return time.Time{}, fmt.Errorf("failed to parse JWT claims: %w", err)
	}
	if claims.Exp == 0 {
		return time.Time{}, fmt.Errorf("JWT missing exp claim")
	}

	return time.Unix(claims.Exp, 0), nil
}

