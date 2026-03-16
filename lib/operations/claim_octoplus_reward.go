package operations

type ClaimOctoplusRewardVariables struct {
	AccountNumber string `json:"accountNumber"`
	OfferSlug     string `json:"offerSlug"`
}

type ClaimOctoplusRewardPayload struct {
	RewardID FlexibleInt `json:"rewardId"`
}

type ClaimOctoplusRewardData struct {
	ClaimOctoplusReward ClaimOctoplusRewardPayload `json:"claimOctoplusReward"`
}

func NewClaimOctoplusRewardRequest(accountNumber, offerSlug string) GraphQLRequest[ClaimOctoplusRewardVariables] {
	return GraphQLRequest[ClaimOctoplusRewardVariables]{
		OperationName: "claimOctoplusReward",
		Variables: ClaimOctoplusRewardVariables{
			AccountNumber: accountNumber,
			OfferSlug:     offerSlug,
		},
		Query: `mutation claimOctoplusReward($accountNumber: String!, $offerSlug: String!) {
			claimOctoplusReward(accountNumber: $accountNumber, offerSlug: $offerSlug) {
				rewardId
			}
		}`,
	}
}
