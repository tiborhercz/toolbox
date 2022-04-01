# cli-toolbox

A collection of various tools like base64 encoding, jwt decoding, ipv4 cidr calculation, and more.
The tools can be accessed from the command line or by using the web UI.

## Installation

Install instructions

### Binary

Get the binary from the [GitHub release page](https://github.com/tiborhercz/cli-toolbox/releases)

### Brew

```shell
brew tap tiborhercz/cli-toolbox
brew install cli-toolbox
```

### Compile

Compiling the binary yourself is possible, either manually or by using the make command.

### Manual
Steps to compile:
1. Build the Wasm binary
```shell
GOOS=js GOARCH=wasm go build -o static/UI/public/main.wasm wasm/*.go
```
2. Build the Vue.js frontend. Run `npm run build` inside the `static/UI/public/` directory
3. Run `go build .` in the root directory
4. Run cli-toolbox with `./cli-toolbox`

### Makefile
To build the cli-toolbox run: `make build`

## Usage

```
cli-toolbox [command]

Available Commands:
  base64      Encode and decode base64 strings
  cidr        Calculate IPv4 CIDR ranges
  completion  generate the autocompletion script for the specified shell
  hash        Hash
  help        Help about any command
  jwtdecode   Decode jwt token
  webui       Launch the web UI for the Toolbox
```

### Launching the web UI
```shell
cli-toolbox webui
```
