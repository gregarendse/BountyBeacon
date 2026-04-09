package client

import "github.com/gregarendse/BountyBeacon/lib/operations"

type GraphQLRequest[TVariables any] = operations.GraphQLRequest[TVariables]
type GraphQLError = operations.GraphQLError
type GraphQLResponse[TData any] = operations.GraphQLResponse[TData]

type LoggedInUserForBreadcrumbsVariables = operations.LoggedInUserForBreadcrumbsVariables
type LoggedInUserForBreadcrumbsData = operations.LoggedInUserForBreadcrumbsData
type Account = operations.Account

type OctoplusRewardsVariables = operations.OctoplusRewardsVariables
type OctoplusRewardsData = operations.OctoplusRewardsData
type OctoplusReward = operations.OctoplusReward
type Voucher = operations.Voucher
type OfferSummary = operations.OfferSummary

type OctoplusOfferBySlugVariables = operations.OctoplusOfferBySlugVariables
type OctoplusOfferBySlugData = operations.OctoplusOfferBySlugData
type OctoplusOffer = operations.OctoplusOffer

type ClaimOctoplusRewardVariables = operations.ClaimOctoplusRewardVariables
type ClaimOctoplusRewardPayload = operations.ClaimOctoplusRewardPayload
type ClaimOctoplusRewardData = operations.ClaimOctoplusRewardData

type OctoplusRewardsByIDVariables = operations.OctoplusRewardsByIDVariables

type ClaimAbility = operations.ClaimAbility
type FlexibleInt = operations.FlexibleInt

type ObtainKrakenTokenVariables = operations.ObtainKrakenTokenVariables
type ObtainKrakenTokenData = operations.ObtainKrakenTokenData

