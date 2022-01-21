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

type CidrOptions struct {
	IpAddress         string
	PrefixLength      int64
	SubnetMaskAddress string
}

type PasswordOptions struct {
	Password  string
	Algorithm string
	Cost      int
}

type HashOptions struct {
	Algorithm string
}
