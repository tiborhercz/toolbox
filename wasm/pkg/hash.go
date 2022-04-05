package pkg

import (
	"fmt"
	"github.com/tiborhercz/toolbox/pkg/hash"
	"syscall/js"
)

func GetSupportedHashingAlgorithms(this js.Value, args []js.Value) any {
	var algorithms = hash.GetSupportedHashingAlgorithms()

	a := make([]interface{}, len(algorithms))
	for i, v := range algorithms {
		a[i] = v
	}

	return a
}

func Hash(this js.Value, args []js.Value) any {
	hash, err := hash.Execute([]byte(args[0].String()), args[1].String())
	if err != nil {
		return fmt.Sprint(err)
	}

	return hash
}
