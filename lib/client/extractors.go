package client

import "fmt"

func ExtractRewardStatus(result *GraphQLResponse[OctoplusRewardsData]) (string, error) {
	if len(result.Errors) > 0 {
		return "", fmt.Errorf("graphql error: %s", result.Errors[0].Message)
	}
	rewards := result.Data.OctoplusRewards
	if len(rewards) == 0 {
		return "", fmt.Errorf("missing octoplusRewards entries")
	}
	status := rewards[0].Status
	if status == "" {
		return "", fmt.Errorf("missing status field")
	}
	return status, nil
}

func ExtractClaimRewardID(result *GraphQLResponse[ClaimOctoplusRewardData]) (int, error) {
	if len(result.Errors) > 0 {
		return 0, fmt.Errorf("graphql error: %s", result.Errors[0].Message)
	}
	rewardID := int(result.Data.ClaimOctoplusReward.RewardID)
	if rewardID <= 0 {
		return 0, fmt.Errorf("missing rewardId field")
	}
	return rewardID, nil
}

func ExtractClaimAbility(result *GraphQLResponse[OctoplusOfferBySlugData]) (bool, string, error) {
	if len(result.Errors) > 0 {
		return false, "", fmt.Errorf("graphql error: %s", result.Errors[0].Message)
	}
	offer := result.Data.OctoplusOffer
	if offer.Slug == "" {
		return false, "", fmt.Errorf("missing octoplusOffer field")
	}
	claimAbility := offer.ClaimAbility
	return claimAbility.CanClaimOffer, claimAbility.CannotClaimReason, nil
}
