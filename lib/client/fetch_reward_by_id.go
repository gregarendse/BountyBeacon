package client

import (
	"context"

	"github.com/gregarendse/BountyBeacon/lib/operations"
)

func (client *Client) FetchRewardByID(ctx context.Context, rewardID int) (*GraphQLResponse[OctoplusRewardsData], error) {
	return MakeRequest[OctoplusRewardsByIDVariables, OctoplusRewardsData](ctx, client.AccessToken, operations.NewOctoplusRewardsByIDRequest(client.AccountID, rewardID))
}
