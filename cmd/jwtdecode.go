package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/jwtdecode"
	"github.com/tiborhercz/cli-toolbox/internal/model"
)

var (
	jwtOptions model.JwtOptions

	jwtdecodeCmd = &cobra.Command{
		Use:   "jwtdecode",
		Short: "Decode jwt token",
		Run: func(cmd *cobra.Command, args []string) {
			jwtdecode.Process(jwtOptions.Value)
		},
	}
)

func init() {
	rootCmd.AddCommand(jwtdecodeCmd)
	jwtdecodeCmd.Flags().StringVarP(&jwtOptions.Value, "value", "v", "", "Value string")
}
