package client

import (
	"context"
	"fmt"

	"github.com/gregarendse/BountyBeacon/lib/operations"
)

func (client *Client) FetchAccountID(ctx context.Context) error {
	result, err := MakeRequest[LoggedInUserForBreadcrumbsVariables, LoggedInUserForBreadcrumbsData](ctx, client.AccessToken, operations.NewLoggedInUserForBreadcrumbsRequest(), ApiOctopusEndpoint)
	if err != nil {
		return err
	}

	if len(result.Errors) > 0 {
		return fmt.Errorf("fetch account request failed: %s", result.Errors[0].Message)
	}

	accounts := result.Data.Viewer.Accounts
	if len(accounts) > 0 {
		client.AccountID = accounts[0].Number
		return nil
	}
	return fmt.Errorf("no accounts found")
}
