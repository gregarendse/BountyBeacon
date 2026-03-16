package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func claimCommand(ctx *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim",
		Short: "Claim an offer and poll for ISSUED reward status",
		RunE: func(cmd *cobra.Command, args []string) error {
			pollCfg, err := resolveClaimPollConfig()
			if err != nil {
				return fmt.Errorf("invalid claim options: %w", err)
			}

			offerSlug := strings.TrimSpace(viper.GetString("claim_offer"))
			if offerSlug == "" {
				return fmt.Errorf("invalid --offer value")
			}

			return handleClaim(cmd.Context(), ctx, offerSlug, pollCfg)
		},
	}

	cmd.Flags().String("offer", "caffe-nero", "Offer slug to claim")
	cmd.Flags().String("claim-poll-interval", defaultClaimPollInterval.String(), "Polling interval while waiting for reward issuance")
	cmd.Flags().String("claim-timeout", defaultClaimPollTimeout.String(), "Timeout while waiting for reward issuance")
	_ = viper.BindPFlag("claim_offer", cmd.Flags().Lookup("offer"))
	_ = viper.BindPFlag("claim_poll_interval", cmd.Flags().Lookup("claim-poll-interval"))
	_ = viper.BindPFlag("claim_poll_timeout", cmd.Flags().Lookup("claim-timeout"))
	return cmd
}

func handleClaim(runCtx context.Context, ctx *commandContext, offerSlug string, pollCfg claimPollConfig) error {
	if err := claimOfferAndPrintReward(runCtx, ctx.client, offerSlug, pollCfg); err != nil {
		return fmt.Errorf("claim failed: %w", err)
	}
	return nil
}
