package client

import (
	"context"
	"time"
)

type OctopusClient interface {
	RenewToken(ctx context.Context) error
	FetchAccountID(ctx context.Context) error
	EnsureFreshToken(ctx context.Context, refreshLeadTime time.Duration) error
	FetchRewards(ctx context.Context) (*GraphQLResponse[OctoplusRewardsData], error)
	FetchOfferBySlug(ctx context.Context, offerSlug string) (*GraphQLResponse[OctoplusOfferBySlugData], error)
	ClaimOfferBySlug(ctx context.Context, offerSlug string) (*GraphQLResponse[ClaimOctoplusRewardData], error)
	FetchRewardByID(ctx context.Context, rewardID int) (*GraphQLResponse[OctoplusRewardsData], error)
	WaitForRewardIssued(ctx context.Context, rewardID int, timeout, interval time.Duration) (*GraphQLResponse[OctoplusRewardsData], error)
}
