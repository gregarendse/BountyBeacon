package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func rewardsCommand(ctx *commandContext) *cobra.Command {
	return &cobra.Command{
		Use:   "rewards",
		Short: "List existing Octoplus rewards",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := ctx.client.FetchRewards(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to fetch rewards: %w", err)
			}
			printJSON(result)
			return nil
		},
	}
}
