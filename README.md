# Toolbox

A collection of various tools like encoding, decoding, hashing, ipv4 cidr calculation, and more.
The tools can be accessed from the command line or by using the web UI.

A lot of tools like these are hosted online and use backend system to do the processing. 
This means you don't know what they do with your sensitive data. Another downside of online tools is that they are littered with ads.

This toolbox runs locally on your machine and all the processing is done in your browser using WebAssembly or JavaScript.

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

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
