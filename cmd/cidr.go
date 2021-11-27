package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/cidr"
)

var (
	cidrOptions model.CidrOptions

	cidrCmd = &cobra.Command{
		Use:   "cidr",
		Short: "Calculate IPv4 CIDR ranges",
		Run: func(cmd *cobra.Command, args []string) {
			cidr.Main(cidrOptions.IpAddress, cidrOptions.PrefixLength)
		},
	}
)

func init() {
	rootCmd.AddCommand(cidrCmd)
	cidrCmd.Flags().StringVarP(&cidrOptions.IpAddress, "ipaddress", "i", "", "ip address. Example inputs: 10.0.0.0/16")
}
