package operations

type OctoplusOfferBySlugVariables struct {
	AccountNumber string `json:"accountNumber"`
	Slug          string `json:"slug"`
}

type ClaimAbility struct {
	CanClaimOffer     bool   `json:"canClaimOffer"`
	CannotClaimReason string `json:"cannotClaimReason"`
}

type OctoplusOffer struct {
	Slug         string       `json:"slug"`
	Name         string       `json:"name"`
	ClaimAbility ClaimAbility `json:"claimAbility"`
	ClaimBy      string       `json:"claimBy"`
}

type OctoplusOfferBySlugData struct {
	OctoplusOffer OctoplusOffer `json:"octoplusOffer"`
}

func NewOctoplusOfferBySlugRequest(accountNumber, slug string) GraphQLRequest[OctoplusOfferBySlugVariables] {
	return GraphQLRequest[OctoplusOfferBySlugVariables]{
		OperationName: "getOctoplusOfferBySlug",
		Variables: OctoplusOfferBySlugVariables{
			AccountNumber: accountNumber,
			Slug:          slug,
		},
		Query: `query getOctoplusOfferBySlug($accountNumber: String!, $slug: String!) {
			octoplusOffer(accountNumber: $accountNumber, slug: $slug) {
				slug
				name
				claimAbility {
					canClaimOffer
					cannotClaimReason
				}
				claimBy
			}
		}`,
	}
}
