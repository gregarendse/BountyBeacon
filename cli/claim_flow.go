package cli

import (
	"context"
	"fmt"

	libclient "github.com/gregarendse/BountyBeacon/lib/client"
)

func claimOfferAndPrintReward(ctx context.Context, client *libclient.Client, offerSlug string, pollCfg claimPollConfig) error {
	claimResult, err := client.ClaimOfferBySlug(ctx, offerSlug)
	if err != nil {
		return fmt.Errorf("claim failed: %w", err)
	}

	logInfo("claim submitted", "offer", offerSlug)
	printJSON(claimResult)

	rewardID, err := libclient.ExtractClaimRewardID(claimResult)
	if err != nil {
		logWarn("claim submitted but reward id missing", "offer", offerSlug, "error", err)
		return nil
	}

	rewardResult, err := client.WaitForRewardIssued(ctx, rewardID, pollCfg.timeout, pollCfg.interval)
	if err != nil {
		logWarn("claim succeeded but reward not issued before timeout", "offer", offerSlug, "reward_id", rewardID, "error", err)
		return nil
	}

	logInfo("reward issued", "offer", offerSlug, "reward_id", rewardID)
	printJSON(rewardResult)
	return nil
}
