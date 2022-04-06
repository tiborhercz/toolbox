package pkg

import (
	"fmt"
	"github.com/tiborhercz/toolbox/pkg/base64"
	"syscall/js"
)

func ProcessBase64(this js.Value, args []js.Value) interface{} {
	var result, err = base64.Process(args[0].String(), args[1].Bool(), args[2].Bool())
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return result
}
