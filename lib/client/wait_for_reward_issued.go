package client

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidWaitTimeout  = errors.New("invalid wait timeout")
	ErrInvalidWaitInterval = errors.New("invalid wait interval")
)

func (client *Client) WaitForRewardIssued(ctx context.Context, rewardID int, timeout, interval time.Duration) (*GraphQLResponse[OctoplusRewardsData], error) {
	if timeout <= 0 {
		return nil, fmt.Errorf("%w: timeout must be greater than zero", ErrInvalidWaitTimeout)
	}
	if interval <= 0 {
		return nil, fmt.Errorf("%w: interval must be greater than zero", ErrInvalidWaitInterval)
	}

	waitCtx, cancel := context.WithTimeoutCause(ctx, timeout, fmt.Errorf("timed out waiting for ISSUED status"))
	defer cancel()

	ticks := time.Tick(interval)
	lastStatus := ""
	for {
		if err := client.EnsureFreshToken(waitCtx, defaultTokenRefreshLeadTime); err != nil {
			return nil, fmt.Errorf("token refresh failed while polling reward: %w", err)
		}

		rewardResult, err := client.FetchRewardByID(waitCtx, rewardID)
		if err != nil {
			select {
			case <-waitCtx.Done():
				return nil, fmt.Errorf("%w: %w", context.Cause(waitCtx), err)
			case <-ticks:
				continue
			}
		}

		status, parseErr := ExtractRewardStatus(rewardResult)
		if parseErr == nil && status == "ISSUED" {
			return rewardResult, nil
		}
		if parseErr == nil {
			lastStatus = status
		}

		select {
		case <-waitCtx.Done():
			if parseErr != nil {
				return nil, fmt.Errorf("unable to parse reward status: %w", parseErr)
			}
			if lastStatus != "" {
				return nil, fmt.Errorf("%w (last status=%q)", context.Cause(waitCtx), lastStatus)
			}
			return nil, context.Cause(waitCtx)
		case <-ticks:
		}
	}
}
