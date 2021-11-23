package jwtdecode

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func Process(value string) {
	split := strings.Split(value, ".")

	if len(split) != 3 {
		log.Fatal("JWT token should split up by three dots")
	}

	for _, value := range split[0:2] {
		decodedString := decodeB64(value)
		output(decodedString)
	}
}

func decodeB64(value string) []byte {
	decodedString, err := b64.RawStdEncoding.DecodeString(value)

	if err != nil {
		log.Fatal(err)
	}

	if !json.Valid(decodedString){
		log.Fatal("invalid JSON")
	}

	return decodedString
}

func output(value []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, value, "", "\t")

	fmt.Println(prettyJSON.String())

	if err != nil {
		log.Fatal(err)
	}
}
