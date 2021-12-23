package jwtdecode

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func Process(value string) (string, error) {
	split := strings.Split(value, ".")

	if len(split) != 3 {
		return "", errors.New("JWT token should split up by three dots")
	}

	var (
		errorValue  error
		returnValue strings.Builder
	)

	for _, value := range split[0:2] {
		decodedString, err := decodeB64(value)
		if err != nil {
			errorValue = err
			break
		}

		if !json.Valid(decodedString) {
			errorValue = errors.New("invalid JSON")
			break
		}

		prettyJson, err := prettifyJson(decodedString)
		if err != nil {
			errorValue = err
			break
		}

		returnValue.WriteString(prettyJson)
	}

	if errorValue != nil {
		return "", errorValue
	}

	return returnValue.String(), nil
}

func decodeB64(value string) ([]byte, error) {
	decodedString, err := b64.RawStdEncoding.DecodeString(value)

	if err != nil {
		return nil, err
	}

	return decodedString, nil
}

func prettifyJson(value []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, value, "", "\t")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
