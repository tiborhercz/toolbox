package model

type Base64Options struct {
	Value       string
	Path        string
	Decode      bool
	Urlencoding bool
}

type JwtOptions struct {
	Value       string
}

type cidrOptions struct {
	address       string
}
