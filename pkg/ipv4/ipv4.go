package ipv4

import (
	"math"
)

func GetTotalCidrIpAddresses(ipPrefixNumber byte) int {
	return calculateTotalIpAddresses(ipPrefixNumber)
}

// calculateTotalIpAddresses Subtract the number of network bits from 32. Raise 2 to that power.
func calculateTotalIpAddresses(ipPrefixNumber byte) int {
	return int(math.Pow(2, float64(32 - ipPrefixNumber)))
}
