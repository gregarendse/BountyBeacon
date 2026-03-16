package operations

type OctoplusRewardsVariables struct {
	AccountNumber string `json:"accountNumber"`
}

func NewOctoplusRewardsRequest(accountNumber string) GraphQLRequest[OctoplusRewardsVariables] {
	return GraphQLRequest[OctoplusRewardsVariables]{
		OperationName: "getOctoplusRewards",
		Variables:     OctoplusRewardsVariables{AccountNumber: accountNumber},
		Query: `query getOctoplusRewards($accountNumber: String!) {
			octoplusRewards(accountNumber: $accountNumber) {
				id
				claimedAt
				vouchers {
					__typename
					... on OctoplusVoucherType {
						code
						expiresAt
					}
					... on ShoptopusVoucherType {
						code
					}
				}
				offer {
					name
					slug
				}
			}
		}`,
	}
}
