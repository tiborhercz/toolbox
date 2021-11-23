package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"os"
)

var (
	options model.Options

	rootCmd = &cobra.Command{
		Use:   "cli-toolbox",
		Short: "cli toolbox",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
