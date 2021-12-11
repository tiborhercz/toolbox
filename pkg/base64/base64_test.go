package base64

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	value := "This is a test value"
	test := "VGhpcyBpcyBhIHRlc3QgdmFsdWU="
	encodedValue := encode(value, false)

	if encodedValue != test {
		t.Fatalf("Encoded value should be 'VGhpcyBpcyBhIHRlc3QgdmFsdWU=' instead is %v", encodedValue)
	}
}

func TestBase64Decode(t *testing.T) {
	value := "VGhpcyBpcyBhIHRlc3QgdmFsdWU="
	test := "This is a test value"
	decodedValue, _ := decode(value)

	if decodedValue != test {
		t.Fatalf("Encoded value should be 'VGhpcyBpcyBhIHRlc3QgdmFsdWU=' instead is %v", decodedValue)
	}
}
