package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/toolbox/internal/model"
	"github.com/tiborhercz/toolbox/pkg/hash"
	"strings"
)

var (
	hashOptions model.HashOptions

	hashCmd = &cobra.Command{
		Use:   "hash",
		Short: "Hash",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || args[0] == "" {
				cmd.Help()
				fmt.Print("\n")
				logrus.Fatal("must provide an argument")
			}

			algorithm := strings.ToUpper(hashOptions.Algorithm)
			hashData, err := hash.Execute([]byte(args[0]), algorithm)
			if err != nil {
				logrus.Fatal(err)
			}

			logrus.Infof("Hash algorithm: %v\n%v", algorithm, hashData)
		},
	}

	getSupportedHashingAlgorithmsCmd = &cobra.Command{
		Use:   "supported",
		Short: "Get supported hashing algorithms",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("Supported hashing algorithms: \n%v", strings.Join(hash.GetSupportedHashingAlgorithms(), "\n"))
		},
	}
)

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.AddCommand(getSupportedHashingAlgorithmsCmd)

	hashCmd.Flags().StringVarP(&hashOptions.Algorithm,
		"algorithm",
		"a",
		"SHA256",
		"Algorithm. Supported algorithms: "+strings.Join(hash.GetSupportedHashingAlgorithms(), " | "))
}
