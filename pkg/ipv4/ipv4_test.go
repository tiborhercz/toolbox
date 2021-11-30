package ipv4

import "testing"

func TestCalculateTotalIpAddresses(t *testing.T) {
	value := byte(10)
	test := 4194304
	totalIpAddresses := calculateTotalIpAddresses(value)

	if totalIpAddresses != test {
		t.Fatalf("totalIpAddresses is %v expected %v", totalIpAddresses, test)
	}
}