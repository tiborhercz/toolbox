package ipv6

import (
	"math/big"
	"testing"
)

func TestCalculateTotalIpAddresses(t *testing.T) {
	value := int64(10)
	test := int64(18014398509481984)
	totalIpAddresses := calculateTotalIpAddresses(64, value)

	if totalIpAddresses.Cmp(big.NewInt(test)) != 0  {
		t.Fatalf("totalIpAddresses is %v expected %v", totalIpAddresses, test)
	}
}