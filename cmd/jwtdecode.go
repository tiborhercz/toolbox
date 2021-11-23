package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/jwtdecode"
)

var (
	jwtdecodeCmd = &cobra.Command{
		Use:   "jwtdecode",
		Short: "Decode jwt token",
		Run: func(cmd *cobra.Command, args []string) {
			jwtdecode.Process(options.Value)
		},
	}
)

func init() {
	rootCmd.AddCommand(jwtdecodeCmd)
	jwtdecodeCmd.Flags().StringVarP(&options.Value, "value", "v", "", "Value string")
}
