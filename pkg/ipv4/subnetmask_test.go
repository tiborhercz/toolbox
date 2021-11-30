package ipv4

import (
	"bytes"
	"testing"
)

func TestCidrNumberToSubnetMask(t *testing.T) {
	subnetMask := []byte{255, 255, 0, 0}
	testSubnetMask := cidrNumberToSubnetMask(16)

	if !bytes.Equal(subnetMask, testSubnetMask) {
		t.Fatalf("subnetMask is %v expected %v", testSubnetMask, subnetMask)
	}
}
