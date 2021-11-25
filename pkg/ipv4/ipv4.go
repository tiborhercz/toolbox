package ipv4

import (
	"math"
)

func GetTotalCidrIpAddresses(networkSize int64) int {
	return calculateTotalIpAddresses(networkSize)
}

// Subtract the number of network bits from 32. Raise 2 to that power.
func calculateTotalIpAddresses(networkSize int64) int {
	return int(math.Pow(2, float64(32 - networkSize)))
}
