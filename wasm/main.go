package main

import (
	"github.com/tiborhercz/cli-toolbox/wasm/pkg"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("wasmBase64", js.FuncOf(pkg.ProcessBase64))
	global.Set("wasmJwtDecode", js.FuncOf(pkg.ProcessJwtDecode))
	global.Set("wasmIpv4", js.FuncOf(pkg.ProcessIpv4))
	<-done
}
