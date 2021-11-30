package cidr

import (
	"github.com/sirupsen/logrus"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/ipv4"
	"net"
	"strconv"
	"strings"
)

var (
	IPv4Data = model.IPv4OutputData {}
)

func IpAddress(ipAddress string) {
	var (
		parsedIpAddress net.IP
		parsedIpNetAddress net.IPNet
		cidrNumber byte
	)

	parsedIpAddress, parsedIpNetAddress = parseIpAddress(ipAddress)
	cidrNumber = getCidrNumberFromIp(parsedIpNetAddress.String())
	IPv4Data.CidrRange = parsedIpNetAddress.String()


	if parsedIpAddress != nil {
		processIpv4(cidrNumber, parsedIpAddress)

		output()
	}

	if parsedIpAddress.To4() == nil && parsedIpAddress.To16() != nil {
		logrus.Println("IPv6 is not yet supported")
	}
}

func SubnetMask(subnetMaskAddress string) {
	var (
		parsedIpMask net.IPMask
	)

	parsedIpMask = parseIpMask(subnetMaskAddress)

	cidr, _ := parsedIpMask.Size()

	if (parsedIpMask == nil) || (cidr == 0 && strings.Compare("0.0.0.0", subnetMaskAddress) != 0) {
		logrus.Fatalf("%s doesn't look like a valid IPv4 netmask", subnetMaskAddress)
	}

	processIpv4(byte(cidr), nil)

	output()
}

func output() {
	if IPv4Data.CidrRange != "" {
		logrus.Printf("CIDR range: %v \n", IPv4Data.CidrRange)
	}

	if IPv4Data.SubnetMask != "" {
		logrus.Printf("Subnetmask: %v \n", IPv4Data.SubnetMask)
	}

	if IPv4Data.FirstIp != "" {
		logrus.Printf("First IP: %v \n", IPv4Data.FirstIp)
	}

	if IPv4Data.LastIp != "" {
		logrus.Printf("Last IP: %v \n", IPv4Data.LastIp)
	}

	if IPv4Data.TotalIpAddresses != "" {
		logrus.Printf("Total ipv4 addresses: %v \n", IPv4Data.TotalIpAddresses)
	}
}

func parseIpAddress(ipAddress string) (net.IP, net.IPNet) {
	parsedIp, parsedIpNet, err := net.ParseCIDR(ipAddress)

	if err != nil {
		logrus.Fatalln(err)
	}

	return parsedIp, *parsedIpNet
}

func parseIpMask(subnetMask string) net.IPMask {
	parsedSubnetMask := net.IPMask(net.ParseIP(subnetMask).To4())

	return parsedSubnetMask
}

func getCidrNumberFromIp(ipNetAddress string) byte {
	ipAddressCidrPrefix := strings.Split(ipNetAddress, "/")[1]
	networkSize, _ := strconv.ParseInt(ipAddressCidrPrefix, 10, 10)

	return byte(networkSize)
}

func processIpv4(cidrNumber byte, ipAddress net.IP) {
	IPv4Data.TotalIpAddresses = strconv.Itoa(ipv4.GetTotalCidrIpAddresses(cidrNumber))
	IPv4Data.SubnetMask = ipv4.GetSubnetMask(cidrNumber)

	if ipAddress != nil && cidrNumber >= 8 {
		IPv4Data.FirstIp, IPv4Data.LastIp = ipv4.GetFirstLastIp(ipAddress, cidrNumber)
	}
}

//func processIpv6(ipAddress net.IP, ipNetAddress net.IPNet, cidrNumber int64, ipCidrPrefix int64) {
//	if cidrNumber > ipCidrPrefix {
//		logrus.Fatalf("The IPs cidr prefix is bigger then the set cidr prefix option. Cidr prefix from IP is: %v. Set cidr prefix is %v", cidrNumber, ipCidrPrefix)
//	}
//
//	TotalIpAddressesBigInt := ipv6.GetTotalIpAddresses(cidrNumber, ipCidrPrefix)
//	IPv6Data.TotalIpAddresses = TotalIpAddressesBigInt.String()
//}
