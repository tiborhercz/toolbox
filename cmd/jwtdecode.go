package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/jwtdecode"
	"log"
)

var (
	jwtOptions model.JwtOptions

	jwtdecodeCmd = &cobra.Command{
		Use:   "jwtdecode",
		Short: "Decode jwt token",
		Run: func(cmd *cobra.Command, args []string) {
			jwtString, err := jwtdecode.Process(jwtOptions.Value)
			if err != nil {
				log.Fatal(err)
			}

			logrus.Info("\n" + jwtString)
		},
	}
)

func init() {
	rootCmd.AddCommand(jwtdecodeCmd)
	jwtdecodeCmd.Flags().StringVarP(&jwtOptions.Value, "value", "v", "", "Value string")
}
