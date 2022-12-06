package jsonyaml

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/tiborhercz/toolbox/internal/utils"
)

func JsonToYaml(json []byte) {
	y, err := yaml.JSONToYAML(json)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))

	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
}

func YamlToJson(yamlContent []byte) {
	json, err := yaml.YAMLToJSON(yamlContent)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(json))
	fmt.Println(utils.PrettifyJson(json))
}
