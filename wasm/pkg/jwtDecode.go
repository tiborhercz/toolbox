package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiborhercz/cli-toolbox/pkg/jwtdecode"
	"syscall/js"
)

type JwtResponse struct {
	Header  string `json:"header"`
	Payload string `json:"payload"`
}

func ProcessJwtDecode(this js.Value, args []js.Value) any {
	jwtData, err := jwtdecode.Process(args[0].String())
	if err != nil {
		return ""
	}

	jwtHeader, err := prettifyJson(jwtData[0])
	if err != nil {
		return ""
	}

	jwtPayload, err := prettifyJson(jwtData[1])
	if err != nil {
		return ""
	}

	JwtResponse := JwtResponse{
		Header:  jwtHeader,
		Payload: jwtPayload,
	}

	jsonData, err := json.Marshal(JwtResponse)

	if err != nil {
		fmt.Println("Unable to convert the struct to a JSON string")
	} else {
		fmt.Println(string(jsonData))
	}

	return string(jsonData)
}

func prettifyJson(value []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, value, "", "\t")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
