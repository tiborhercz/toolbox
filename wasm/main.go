package main

import (
	"fmt"
	"github.com/tiborhercz/cli-toolbox/pkg/base64"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("wasmBase64", js.FuncOf(processBase64))
	global.Set("wasmJwtDecode", js.FuncOf(processJwtDecode))
	global.Set("wasmIpv4Cidr", js.FuncOf(processIpv4))
	<-done
}

func processBase64(this js.Value, args []js.Value) any {
	var result, err = base64.Process(args[0].String(), args[1].Bool(), args[2].Bool())
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return result
}
