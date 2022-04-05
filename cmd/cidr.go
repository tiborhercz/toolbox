package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/toolbox/internal/cidr"
)

var (
	cidrCmd = &cobra.Command{
		Use:   "cidr",
		Short: "Calculate IPv4 CIDR ranges",
	}

	cidrIpAddressCmd = &cobra.Command{
		Use:   "ip",
		Short: "Calculate IPv4 CIDR ranges from IP address",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Fatal("requires an argument. Example argument: 10.0.0.0/16")
			}

			cidr.IpAddress(args[0])
		},
	}

	cidrSubnetMaskCmd = &cobra.Command{
		Use:   "subnetmask",
		Short: "Calculate IPv4 CIDR ranges from subnetmask",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Fatal("requires an argument. Example argument: 255.255.255.0")
			}

			cidr.SubnetMask(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(cidrCmd)
	cidrCmd.AddCommand(cidrIpAddressCmd, cidrSubnetMaskCmd)
}
