package main

import (
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/jwt"
	"github.com/tiborhercz/cli-toolbox/internal/model"
)

var (
	options model.Options

	rootCmd = &cobra.Command{
		Use:   "jwt",
		Short: "Encode and decode jwt strings",
		Run: func(cmd *cobra.Command, args []string) {
			jwt.Process(options.Value)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&options.Value, "value", "v", "", "Value string")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
