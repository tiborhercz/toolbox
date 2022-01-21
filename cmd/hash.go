package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/hash"
	"log"
)

var (
	hashOptions model.HashOptions

	hashCmd = &cobra.Command{
		Use:   "hash",
		Short: "Hash",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || args[0] == "" {
				logrus.Fatal("must provide an argument")
			}

			hashData, err := hash.Hash(hashOptions.Algorithm)
			if err != nil {
				log.Fatal(err)
			}

			logrus.Infof("%v", hashData)
		},
	}
)

func init() {
	rootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&hashOptions.Algorithm, "algorithm", "a", "bcrypt", "Algorithm. Supported algorithms: ")
}
