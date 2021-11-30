package cidr

import (
	"bytes"
	"encoding/hex"
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

func TestParseIpMask(t *testing.T) {
	value, _ := hex.DecodeString("ffffff00")
	ipMask := parseIpMask("255.255.255.0")

	if !bytes.Equal(value, ipMask) {
		t.Fatalf("ipMask is %v expected %v", ipMask, "ffffff00")
	}
}

func TestGetCidrNumberFromIp(t *testing.T) {
	cidrNumber := getCidrNumberFromIp("192.168.0.0/16")

	if cidrNumber != 16 {
		t.Fatalf("cidrNumber is %v expected %v", cidrNumber, 16)
	}
}
