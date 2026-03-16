package client

import (
	"context"

	"github.com/gregarendse/BountyBeacon/lib/operations"
)

func (client *Client) ClaimOfferBySlug(ctx context.Context, offerSlug string) (*GraphQLResponse[ClaimOctoplusRewardData], error) {
	return MakeRequest[ClaimOctoplusRewardVariables, ClaimOctoplusRewardData](ctx, client.AccessToken, operations.NewClaimOctoplusRewardRequest(client.AccountID, offerSlug))
}
