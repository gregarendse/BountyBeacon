package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	libclient "github.com/gregarendse/BountyBeacon/lib/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func watchCommand(ctx *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Poll offer availability and optionally auto-claim",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleWatch(cmd.Context(), ctx.client)
		},
	}

	cmd.Flags().String("offer", "caffe-nero", "Offer slug to watch")
	cmd.Flags().String("interval", "30s", "Polling interval for offer availability checks")
	cmd.Flags().Bool("auto-claim", false, "Automatically claim when offer is available")
	cmd.Flags().String("claim-poll-interval", defaultClaimPollInterval.String(), "Polling interval while waiting for reward issuance")
	cmd.Flags().String("claim-timeout", defaultClaimPollTimeout.String(), "Timeout while waiting for reward issuance")
	_ = viper.BindPFlag("watch_offer", cmd.Flags().Lookup("offer"))
	_ = viper.BindPFlag("watch_interval", cmd.Flags().Lookup("interval"))
	_ = viper.BindPFlag("watch_auto_claim", cmd.Flags().Lookup("auto-claim"))
	_ = viper.BindPFlag("claim_poll_interval", cmd.Flags().Lookup("claim-poll-interval"))
	_ = viper.BindPFlag("claim_poll_timeout", cmd.Flags().Lookup("claim-timeout"))
	return cmd
}

func handleWatch(ctx context.Context, client *libclient.Client) error {
	interval, err := parseDurationConfig("watch_interval", "watch interval")
	if err != nil {
		return err
	}
	offerSlug := strings.TrimSpace(viper.GetString("watch_offer"))
	if offerSlug == "" {
		return fmt.Errorf("invalid --offer value")
	}
	autoClaim := viper.GetBool("watch_auto_claim")

	pollCfg, err := resolveClaimPollConfig()
	if err != nil {
		return fmt.Errorf("invalid claim poll config: %w", err)
	}

	logInfo("watch started", "offer", offerSlug, "interval", interval.String(), "auto_claim", autoClaim, "claim_poll_interval", pollCfg.interval.String(), "claim_timeout", pollCfg.timeout.String())
	ticks := time.Tick(interval)
	for {
		if err := ctx.Err(); err != nil {
			return context.Cause(ctx)
		}

		if err := client.EnsureFreshToken(ctx, 5*time.Minute); err != nil {
			logWarn("token refresh failed, attempting API key re-login", "offer", offerSlug, "error", err)
			apiKey, credErr := readAPIKey()
			if credErr != nil {
				logWarn("no API key available for re-login", "offer", offerSlug, "error", credErr)
				if err := waitForNextPoll(ctx, ticks); err != nil {
					return err
				}
				continue
			}
			newClient, loginErr := libclient.LoginWithAPIKey(ctx, apiKey)
			if loginErr != nil {
				logWarn("API key re-login failed", "offer", offerSlug, "error", loginErr)
				if err := waitForNextPoll(ctx, ticks); err != nil {
					return err
				}
				continue
			}
			*client = *newClient
			logInfo("API key re-login successful", "offer", offerSlug, "account", client.AccountID)
		}

		checkResult, err := client.FetchOfferBySlug(ctx, offerSlug)
		if err != nil {
			logWarn("offer check failed", "offer", offerSlug, "error", err)
			if err := waitForNextPoll(ctx, ticks); err != nil {
				return err
			}
			continue
		}

		canClaim, reason, err := libclient.ExtractClaimAbility(checkResult)
		if err != nil {
			logWarn("failed to parse offer claimability", "offer", offerSlug, "error", err)
			if err := waitForNextPoll(ctx, ticks); err != nil {
				return err
			}
			continue
		}

		logInfo("offer status", "offer", offerSlug, "can_claim", canClaim, "reason", reason)
		if !canClaim {
			if err := waitForNextPoll(ctx, ticks); err != nil {
				return err
			}
			continue
		}

		logInfo("offer available", "offer", offerSlug)
		if !autoClaim {
			return nil
		}

		if err := claimOfferAndPrintReward(ctx, client, offerSlug, pollCfg); err != nil {
			logWarn("auto-claim failed", "offer", offerSlug, "error", err)
			if err := waitForNextPoll(ctx, ticks); err != nil {
				return err
			}
			continue
		}

		return nil
	}
}

func waitForNextPoll(ctx context.Context, ticks <-chan time.Time) error {
	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	case <-ticks:
		return nil
	}
}
