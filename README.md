# cli-toolbox

A handy collection of cli tools

## Installation

Install instructions

### Binary

Get the binary from the [GitHub release page](https://github.com/tiborhercz/cli-toolbox/releases)

### Brew

```
brew tap tiborhercz/cli-toolbox
brew install cli-toolbox
```

## Usage

```
cli-toolbox [command]
```

### Base64

```
Encode and decode base64 strings

Usage:
  cli-toolbox base64 [flags]

Flags:
  -d, --decode         Decode
  -h, --help           help for base64
  -p, --path string    Path string
  -u, --urlencoding    URLEncoding is the alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.
  -v, --value string   Value string
```

### Jwtdecode

```
Decode jwt token

Usage:
  cli-toolbox jwtdecode [flags]

Flags:
  -h, --help           help for jwtdecode
  -v, --value string   Value string
```

### cidr

```
Calculate IPv4 and IPv6 CIDR ranges

Usage:
  cli-toolbox cidr [flags]

Flags:
  -c, --cidrprefix int     IpCidrPrefix default 64 (default 64)
  -h, --help               help for cidr
  -i, --ipaddress string   ip address
```
