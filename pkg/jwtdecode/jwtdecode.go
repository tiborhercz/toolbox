package jwtdecode

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func Process(value string) ([][]byte, error) {
	split := strings.Split(value, ".")

	if len(split) != 3 {
		return nil, errors.New("JWT token should split up by three dots")
	}

	var (
		errorValue  error
		returnValue [][]byte
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

		returnValue = append(returnValue, decodedString)
	}

	if errorValue != nil {
		return nil, errorValue
	}

	return returnValue, nil
}

func decodeB64(value string) ([]byte, error) {
	decodedString, err := b64.RawStdEncoding.DecodeString(value)

	if err != nil {
		return nil, err
	}

	return decodedString, nil
}
