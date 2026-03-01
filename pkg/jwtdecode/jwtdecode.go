package jwtdecode

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

var ErrInvalidJWTFormat = errors.New("invalid JWT format")

func DecodeRaw(token string) ([]byte, []byte, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, ErrInvalidJWTFormat
	}

	header, err := b64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, err
	}

	payload, err := b64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, nil, err
	}

	return header, payload, nil
}

func Decode(token string) (map[string]any, map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, ErrInvalidJWTFormat
	}

	decode := func(part string) (map[string]any, error) {
		raw, err := b64.RawURLEncoding.DecodeString(part)
		if err != nil {
			return nil, err
		}

		var data map[string]any
		if err := json.Unmarshal(raw, &data); err != nil {
			return nil, err
		}

		return data, nil
	}

	header, err := decode(parts[0])
	if err != nil {
		return nil, nil, err
	}

	payload, err := decode(parts[1])
	if err != nil {
		return nil, nil, err
	}

	return header, payload, nil
}
