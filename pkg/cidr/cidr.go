package cidr

import (
	log "github.com/sirupsen/logrus"
	"github.com/tiborhercz/cli-toolbox/pkg/ipv4"
	"github.com/tiborhercz/cli-toolbox/pkg/ipv6"
	"math/big"
	"net"
	"strconv"
	"strings"
)

func Main(ipAddress string, ipCidrPrefix int64) {
	parsedIpAddress, parsedIpNetAddress := parseIpAddress(ipAddress)

	outputTotalIpAddresses(parsedIpAddress, parsedIpNetAddress, ipCidrPrefix)
}

func outputTotalIpAddresses(ipAddress net.IP, ipNetAddress net.IPNet, ipCidrPrefix int64)  {
	var (
		totalIpAddresses int
		totalIpAddressesBig big.Int
	)

	networkSize := getIpNetworkSize(ipNetAddress.String())

	if ipAddress.To4() != nil {
		totalIpAddresses = ipv4.GetTotalCidrIpAddresses(networkSize)
		log.Printf("Total ipv4 addresses: %d \n", totalIpAddresses)
	} else if ipAddress.To16() != nil {
		if networkSize > ipCidrPrefix {
			log.Fatalf("The IPs cidr prefix is bigger then the set cidr prefix option. Cidr prefix from IP is: %v. Set cidr prefix is %v", networkSize, ipCidrPrefix)
		}

		totalIpAddressesBig = ipv6.GetTotalIpAddresses(networkSize, ipCidrPrefix)
		log.Printf("Total ipv6 addresses: %v \n", totalIpAddressesBig.String())
	}
}

func parseIpAddress(ipAddress string) (net.IP, net.IPNet) {
	parsedIp, parsedIpNet, err := net.ParseCIDR(ipAddress)

	if err != nil {
		log.Fatalln(err)
	}

	return parsedIp, *parsedIpNet
}

func getIpNetworkSize(ipNetAddress string) int64 {
	ipAddressCidrPrefix := strings.Split(ipNetAddress, "/")[1]
	networkSize, _ := strconv.ParseInt(ipAddressCidrPrefix, 10,10)

	return networkSize
}
