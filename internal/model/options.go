package model

type Base64Options struct {
	Value       string
	Path        string
	Decode      bool
	Urlencoding bool
}

type JwtOptions struct {
	Value string
}

type WebUIOptions struct {
	Port string
}

type CidrOptions struct {
	IpAddress         string
	PrefixLength      int64
	SubnetMaskAddress string
}
