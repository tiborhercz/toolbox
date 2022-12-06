package main

import (
	"fmt"
	"github.com/tiborhercz/toolbox/cmd"
	"github.com/tiborhercz/toolbox/internal/logrus"
	"github.com/tiborhercz/toolbox/pkg/conversion/jsonyaml"
	"os"
)

func main() {
	logrus.SetOptions()
	cmd.Execute()

	b, err := os.ReadFile("./test.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	jsonyaml.JsonToYaml(b)

	//yaml123, err := os.ReadFile("./test.yaml") // just pass the file name
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//jsonyaml.YamlToJson(yaml123)
}
