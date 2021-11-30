package ipv4

import (
	"math"
	"net"
)

func GetTotalCidrIpAddresses(cidrNumber byte) int {
	return calculateTotalIpAddresses(cidrNumber)
}

func GetFirstLastIp(ipAddress net.IP, cidrNumber byte) (string, string) {
	firstIp, lastIp := generateFirstLastIp(ipAddress, cidrNumber)

	return net.IP(firstIp).String(), net.IP(lastIp).String()
}

func generateFirstLastIp(ipAddress net.IP, cidrNumber byte) ([]byte, []byte) {
	subnetMask := cidrNumberToSubnetMask(cidrNumber)
	firstIp := make([]byte, 4)
	lastIp := make([]byte, 4)
	octetPos := int(math.Floor(float64(cidrNumber) / float64(8)))

	for i, address := range ipAddress.To4() {
		if i >= octetPos {
			firstIp[i] = 0
			lastIp[i] = 255 - subnetMask[i] // TODO does not work for higher cidrNumber
		} else {
			firstIp[i] = address
			lastIp[i] = address
		}
	}

	return firstIp, lastIp
}

// calculateTotalIpAddresses Subtract the number of network bits from 32. Raise 2 to that power.
func calculateTotalIpAddresses(cidrNumber byte) int {
	return int(math.Pow(2, float64(32 - cidrNumber)))
}
