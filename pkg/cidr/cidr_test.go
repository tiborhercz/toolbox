package cidr

import (
	"testing"
)

func TestIpv4ParseIpAddress(t *testing.T) {
	value := "192.168.0.0/32"
	parsedIp, parsedIpNet := parseIpAddress(value)

	testParsedIp := "192.168.0.0"
	if parsedIp.String() != testParsedIp {
		t.Fatalf("parsedIp is %v expected %v", parsedIp.String(), testParsedIp)
	}

	testParsedIpNet := "192.168.0.0/32"
	if parsedIpNet.String() != testParsedIpNet {
		t.Fatalf("parsedIpNet is %v expected %v", parsedIp.String(), testParsedIpNet)
	}
}

func TestIpv6ParseIpAddress(t *testing.T) {
	value := "2001:db8:3333:4444:5555:6666:7777:8888/128"
	parsedIp, parsedIpNet := parseIpAddress(value)

	testParsedIp := "2001:db8:3333:4444:5555:6666:7777:8888"
	if parsedIp.String() != testParsedIp {
		t.Fatalf("parsedIp is %v expected %v", parsedIp.String(), testParsedIp)
	}

	testParsedIpNet := "2001:db8:3333:4444:5555:6666:7777:8888/128"
	if parsedIpNet.String() != testParsedIpNet {
		t.Fatalf("parsedIpNet is %v expected %v", parsedIp.String(), testParsedIpNet)
	}
}

func TestGetIpNetworkSize(t *testing.T) {
	value := "192.168.0.0/32"
	test := int64(32)
	networkSize := getIpNetworkSize(value)

	if networkSize != test {
		t.Fatalf("networkSize is %v expected %v", networkSize, test)
	}
}
