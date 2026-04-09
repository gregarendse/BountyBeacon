package operations

type ObtainKrakenTokenInput struct {
	APIKey string `json:"APIKey"`
}

type ObtainKrakenTokenVariables struct {
	Input ObtainKrakenTokenInput `json:"input"`
}

type ObtainKrakenTokenResult struct {
	Token            string `json:"token"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn int    `json:"refreshExpiresIn"`
}

type ObtainKrakenTokenData struct {
	ObtainKrakenToken ObtainKrakenTokenResult `json:"obtainKrakenToken"`
}

func NewObtainKrakenTokenRequest(apiKey string) GraphQLRequest[ObtainKrakenTokenVariables] {
	return GraphQLRequest[ObtainKrakenTokenVariables]{
		OperationName: "obtainKrakenToken",
		Variables: ObtainKrakenTokenVariables{
			Input: ObtainKrakenTokenInput{
				APIKey: apiKey,
			},
		},
		Query: `mutation obtainKrakenToken($input: ObtainJSONWebTokenInput!) {
			obtainKrakenToken(input: $input) {
				token
				refreshToken
				refreshExpiresIn
			}
		}`,
	}
}
