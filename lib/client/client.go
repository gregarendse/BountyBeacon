package client

import (
	"time"
)

type Client struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	AccountID    string    `json:"account_id"`
	ClientID     string    `json:"client_id,omitzero"`
	ExpiresAt    time.Time `json:"expires_at,omitzero"`
}

func New() *Client {
	return &Client{}
}

func (client *Client) SetAccessToken(accessToken string) {
	client.AccessToken = accessToken
}
