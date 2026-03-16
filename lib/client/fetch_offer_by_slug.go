package client

import (
	"context"

	"github.com/gregarendse/BountyBeacon/lib/operations"
)

func (client *Client) FetchOfferBySlug(ctx context.Context, offerSlug string) (*GraphQLResponse[OctoplusOfferBySlugData], error) {
	return MakeRequest[OctoplusOfferBySlugVariables, OctoplusOfferBySlugData](ctx, client.AccessToken, operations.NewOctoplusOfferBySlugRequest(client.AccountID, offerSlug))
}
