package client

import (
	"context"

	"github.com/gregarendse/BountyBeacon/lib/operations"
)

func (client *Client) FetchRewards(ctx context.Context) (*GraphQLResponse[OctoplusRewardsData], error) {
	return MakeRequest[OctoplusRewardsVariables, OctoplusRewardsData](ctx, client.AccessToken, operations.NewOctoplusRewardsRequest(client.AccountID))
}
