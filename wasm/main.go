package main

import (
	"github.com/tiborhercz/cli-toolbox/wasm/pkg"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("wasmBase64Process", js.FuncOf(pkg.ProcessBase64))
	global.Set("wasmJwtDecode", js.FuncOf(pkg.ProcessJwtDecode))
	global.Set("wasmIpv4Process", js.FuncOf(pkg.ProcessIpv4))
	global.Set("wasmHashGetSupportedHashingAlgorithms", js.FuncOf(pkg.GetSupportedHashingAlgorithms))
	global.Set("wasmHash", js.FuncOf(pkg.Hash))
	<-done
}
