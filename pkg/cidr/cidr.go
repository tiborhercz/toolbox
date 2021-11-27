package cidr

import (
	log "github.com/sirupsen/logrus"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/ipv4"
	"net"
	"strconv"
	"strings"
)

var (
	IPv4Data = model.IPv4OutputData {}
)

func Main(ipAddress string, ipCidrPrefix int64) {
	parsedIpAddress, parsedIpNetAddress := parseIpAddress(ipAddress)
	ipPrefixNumber := getIpPrefixNumber(parsedIpNetAddress.String())

	if parsedIpAddress.To4() != nil {
		processIpv4(parsedIpAddress, parsedIpNetAddress, ipPrefixNumber)

		outputIPv4()
	} else if parsedIpAddress.To16() != nil {
		log.Println("IPv6 is not yet supported")
	}
}

func outputIPv4() {
	log.Printf("Total ipv4 addresses: %v \n", IPv4Data.TotalIpAddresses)
	log.Printf("Subnetmask: %v \n", IPv4Data.SubnetMask)
}

func parseIpAddress(ipAddress string) (net.IP, net.IPNet) {
	parsedIp, parsedIpNet, err := net.ParseCIDR(ipAddress)

	if err != nil {
		log.Fatalln(err)
	}

	return parsedIp, *parsedIpNet
}

func getIpPrefixNumber(ipNetAddress string) int64 {
	ipAddressCidrPrefix := strings.Split(ipNetAddress, "/")[1]
	networkSize, _ := strconv.ParseInt(ipAddressCidrPrefix, 10, 10)

	return networkSize
}

func processIpv4(ipAddress net.IP, ipNetAddress net.IPNet, ipPrefixNumber int64) {
	IPv4Data.TotalIpAddresses = strconv.Itoa(ipv4.GetTotalCidrIpAddresses(ipPrefixNumber))
	IPv4Data.SubnetMask = ipv4.GetSubnetMask(ipPrefixNumber)
}

//func processIpv6(ipAddress net.IP, ipNetAddress net.IPNet, ipPrefixNumber int64, ipCidrPrefix int64) {
//	if ipPrefixNumber > ipCidrPrefix {
//		log.Fatalf("The IPs cidr prefix is bigger then the set cidr prefix option. Cidr prefix from IP is: %v. Set cidr prefix is %v", ipPrefixNumber, ipCidrPrefix)
//	}
//
//	TotalIpAddressesBigInt := ipv6.GetTotalIpAddresses(ipPrefixNumber, ipCidrPrefix)
//	IPv6Data.TotalIpAddresses = TotalIpAddressesBigInt.String()
//}
