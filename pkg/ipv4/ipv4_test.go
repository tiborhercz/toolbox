package ipv4

import (
	"testing"
)

func TestIpCount(t *testing.T) {
	cidr := byte(15)
	testTotalIpCount := uint32(131072)
	totalIpCount := ipCount(cidr)

	if testTotalIpCount != totalIpCount {
		t.Fatalf("totalIpCount is %v expected %v", totalIpCount, testTotalIpCount)
	}
}

func TestMinIp15(t *testing.T) {
	ip := IP4(0b11000000101010000000000000000000) // 192.168.0.0
	minIp := minIP(ip, 15)

	if ip != minIp {
		t.Fatalf("minIp is %v expected %v", minIp.String(), ip.String())
	}
}

func TestMinIp29(t *testing.T) {
	ip := IP4(0b11000000101010000000101000110010)        // 192.168.10.50
	testMinIp := IP4(0b11000000101010000000101000110000) // 192.168.10.48
	minIp := minIP(ip, 29)

	if testMinIp != minIp {
		t.Fatalf("minIp is %v expected %v", testMinIp.String(), minIp.String())
	}
}

func TestMaxIp15(t *testing.T) {
	ip := IP4(0b11000000101010011111111111111111) // 192.168.255.255
	maxIp := maxIP(ip, 15)

	if ip != maxIp {
		t.Fatalf("maxIp is %v expected %v", maxIp.String(), ip.String())
	}
}

func TestMaxIp29(t *testing.T) {
	ip := IP4(0b11000000101010000000101000110010)        // 192.168.10.50
	testMaxIp := IP4(0b11000000101010000000101000110111) // 192.168.10.55
	maxIp := maxIP(ip, 29)

	if testMaxIp != maxIp {
		t.Fatalf("maxIp is %v expected %v", testMaxIp.String(), maxIp.String())
	}
}

func TestGetMask(t *testing.T) {
	cidrNumber := byte(16)
	testSubnetMask := IP4(0b11111111111111110000000000000000)
	subnetMask := getMask(cidrNumber)

	if testSubnetMask != subnetMask {
		t.Fatalf("maxIp is %v expected %v", subnetMask.String(), testSubnetMask.String())
	}
}
