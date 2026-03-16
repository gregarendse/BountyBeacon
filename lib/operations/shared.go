package operations

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type GraphQLRequest[TVariables any] struct {
	OperationName string     `json:"operationName"`
	Variables     TVariables `json:"variables"`
	Query         string     `json:"query"`
}

type GraphQLError struct {
	Message string         `json:"message"`
	Path    []any          `json:"path,omitzero"`
	Code    string         `json:"code,omitempty"`
	Meta    map[string]any `json:"extensions,omitzero"`
}

type GraphQLResponse[TData any] struct {
	Data   TData          `json:"data"`
	Errors []GraphQLError `json:"errors,omitzero"`
}

type FlexibleInt int

func (f *FlexibleInt) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" || trimmed == "null" {
		*f = 0
		return nil
	}

	if strings.HasPrefix(trimmed, "\"") {
		var raw string
		if err := json.Unmarshal(data, &raw); err != nil {
			return err
		}
		parsed, err := strconv.Atoi(raw)
		if err != nil {
			return fmt.Errorf("invalid integer value %q", raw)
		}
		*f = FlexibleInt(parsed)
		return nil
	}

	var rawInt int
	if err := json.Unmarshal(data, &rawInt); err != nil {
		return err
	}
	*f = FlexibleInt(rawInt)
	return nil
}

type Voucher struct {
	TypeName  string `json:"__typename,omitempty"`
	Code      string `json:"code"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

type OfferSummary struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type OctoplusReward struct {
	ID        int          `json:"id"`
	ClaimedAt string       `json:"claimedAt,omitempty"`
	Status    string       `json:"status,omitempty"`
	Vouchers  []Voucher    `json:"vouchers,omitempty"`
	Offer     OfferSummary `json:"offer"`
}

type OctoplusRewardsData struct {
	OctoplusRewards []OctoplusReward `json:"octoplusRewards"`
}
