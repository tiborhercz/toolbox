package model

type Base64Options struct {
	Path        string
	Decode      bool
	Urlencoding bool
}

type PasswordOptions struct {
	Password       string
	HashedPassword string
	Algorithm      string
	Cost           int
}

type WebUIOptions struct {
	Port string
}

type CidrOptions struct {
	IpAddress         string
	PrefixLength      int64
	SubnetMaskAddress string
}

type HashOptions struct {
	Algorithm string
}
