package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/toolbox/internal/model"
	"github.com/tiborhercz/toolbox/pkg/password"
	"log"
)

var (
	passwordOptions model.PasswordOptions

	passwordCmd = &cobra.Command{
		Use:   "password",
		Short: "Hash and verify passwords",
	}

	passwordHashCmd = &cobra.Command{
		Use:   "hash",
		Short: "hash password",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || args[0] == "" {
				logrus.Fatal("must provide an argument")
			}

			passwordData, err := password.Hash(args[0], passwordOptions.Algorithm, passwordOptions.Cost)
			if err != nil {
				log.Fatal(err)
			}

			logrus.Infof("Algorithm: %v\n%v", passwordOptions.Algorithm, passwordData)
		},
	}

	passwordVerifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "verify password",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || args[0] == "" {
				logrus.Fatal("must provide an argument")
			}

			err := password.Verify(passwordOptions.HashedPassword, args[0])
			if err != nil {
				logrus.Fatal(err)
			}

			logrus.Infof("MATCH - Password: '%v' matched with the hash: %v", args[0], passwordOptions.HashedPassword)
		},
	}
)

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.AddCommand(passwordHashCmd, passwordVerifyCmd)
	passwordVerifyCmd.Flags().StringVarP(&passwordOptions.Algorithm, "algorithm", "a", "bcrypt", "Algorithm. Supported algorithms: "+password.SupportedHashingAlgorithms())
	passwordVerifyCmd.Flags().StringVarP(&passwordOptions.HashedPassword, "hashedPassword", "H", "", "Password hash")
	passwordVerifyCmd.MarkFlagRequired("hashedPassword")

	passwordHashCmd.Flags().StringVarP(&passwordOptions.Algorithm, "algorithm", "a", "bcrypt", "Algorithm. Supported algorithms: "+password.SupportedHashingAlgorithms())
	passwordHashCmd.Flags().IntVarP(&passwordOptions.Cost, "cost", "c", 10, "cost for hashing. Default = 10")
}
