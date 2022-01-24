package model

type Base64Options struct {
	Path        string
	Decode      bool
	Urlencoding bool
}

type PasswordOptions struct {
	Password  string
	Algorithm string
	Cost      int
}

type HashOptions struct {
	Algorithm string
}
