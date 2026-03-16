package operations

type OctoplusRewardsByIDVariables struct {
	AccountNumber string `json:"accountNumber"`
	RewardID      int    `json:"rewardId"`
}

func NewOctoplusRewardsByIDRequest(accountNumber string, rewardID int) GraphQLRequest[OctoplusRewardsByIDVariables] {
	return GraphQLRequest[OctoplusRewardsByIDVariables]{
		OperationName: "getOctoplusRewardsById",
		Variables: OctoplusRewardsByIDVariables{
			AccountNumber: accountNumber,
			RewardID:      rewardID,
		},
		Query: `query getOctoplusRewardsById($accountNumber: String!, $rewardId: Int) {
			octoplusRewards(accountNumber: $accountNumber, rewardId: $rewardId) {
				id
				status
				vouchers {
					... on OctoplusVoucherType {
						code
						expiresAt
					}
					... on ShoptopusVoucherType {
						code
					}
				}
				offer {
					slug
					name
				}
			}
		}`,
	}
}
