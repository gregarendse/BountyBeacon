package cli

import (
	"time"

	libclient "github.com/gregarendse/BountyBeacon/lib/client"
)

const (
	defaultClaimPollInterval = 2 * time.Second
	defaultClaimPollTimeout  = 30 * time.Second
)

type claimPollConfig struct {
	interval time.Duration
	timeout  time.Duration
}

type commandContext struct {
	client *libclient.Client
}
