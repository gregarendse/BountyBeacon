package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func checkCommand(ctx *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check if an offer can be claimed",
		RunE: func(cmd *cobra.Command, args []string) error {
			offerSlug := strings.TrimSpace(viper.GetString("check_offer"))
			if offerSlug == "" {
				return fmt.Errorf("invalid --offer value")
			}

			result, err := ctx.client.FetchOfferBySlug(cmd.Context(), offerSlug)
			if err != nil {
				return fmt.Errorf("failed to check offer: %w", err)
			}
			printJSON(result)
			return nil
		},
	}

	cmd.Flags().String("offer", "caffe-nero", "Offer slug to check")
	_ = viper.BindPFlag("check_offer", cmd.Flags().Lookup("offer"))
	return cmd
}
