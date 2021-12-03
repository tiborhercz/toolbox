package model

type IPv4OutputData struct {
	CidrRange        string
	TotalIpAddresses string
	SubnetMask       string
	FirstIp          string
	LastIp           string
}

type IPv6OutputData struct {
	TotalIpAddresses string
}
