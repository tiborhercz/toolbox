package ipv6

import (
	"math"
	"math/big"
)

func GetTotalIpAddresses(networkSize int64, ipPrefixSize int64) big.Int {
	return calculateTotalIpAddresses(ipPrefixSize, networkSize)
}

// Subtract the number of network bits from 32. Raise 2 to that power.
//func calculateNetworkSize(networkSize int64) big.Int {
//	var i, e = big.NewInt(16), big.NewInt(2)
//	var networkSizeBig big.Int
//	fmt.Println(i.Exp(i, e, nil))
//	return networkSizeBig
//}

func calculateTotalIpAddresses(size int64, mask int64) big.Int {
	ipAddressCount := big.NewInt(0)

	bigInt := big.NewFloat(0)
	bigInt.SetFloat64(math.Pow(2, float64(size-mask)))
	bigIpCountInt, _ := bigInt.Int(nil)

	ipAddressCount.Set(bigIpCountInt)

	return *ipAddressCount
}
