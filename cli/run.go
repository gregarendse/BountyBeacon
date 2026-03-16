package cli

import (
	"fmt"
	"os"
)

func Run(args []string) int {
	configureViperDefaults()
	ctx := &commandContext{}
	root := rootCommand(ctx)
	root.SetArgs(args)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
