package utils

import (
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"strings"
)

func ParseIpAddress(ipAddress string) (net.IP, net.IPNet) {
	parsedIp, parsedIpNet, err := net.ParseCIDR(ipAddress)

	if err != nil {
		logrus.Fatalln(err)
	}

	return parsedIp, *parsedIpNet
}

func GetCidrNumberFromIp(ipNetAddress string) byte {
	ipAddressCidrPrefix := strings.Split(ipNetAddress, "/")[1]
	networkSize, _ := strconv.ParseInt(ipAddressCidrPrefix, 10, 10)

	return byte(networkSize)
}
