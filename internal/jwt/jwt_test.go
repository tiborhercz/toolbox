package jwt

import (
	"bytes"
	"testing"
)

func TestDecodeB64(t *testing.T) {
	value := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	test := []byte{123, 34, 97, 108, 103, 34, 58, 34, 72, 83, 50, 53, 54, 34, 44, 34, 116, 121, 112, 34, 58, 34, 74, 87, 84, 34, 125}
	encodedValue := decodeB64(value)

	if !bytes.Equal(encodedValue, test) {
		t.Fatalf("Encoded value should be '[123 34 97 108 103 34 58 34 72 83 50 53 54 34 44 34 116 121 112 34 58 34 74 87 84 34 125]' instead is %v", encodedValue)
	}
}
