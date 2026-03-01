package jwtdecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DecodeRaw(t *testing.T) {
	tests := []struct {
		name           string
		jwtToken       string
		expectedHeader []byte
		expectedBody   []byte
		expectedErr    error
		wantErr        bool
	}{
		{
			name:           "Valid JWT token",
			jwtToken:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30",
			expectedHeader: []byte(`{"alg":"HS256","typ":"JWT"}`),
			expectedBody:   []byte(`{"sub":"1234567890","name":"John Doe","admin":true,"iat":1516239022}`),
			wantErr:        false,
		},
		{
			name:        "Invalid JWT format",
			jwtToken:    "abc123",
			expectedErr: ErrInvalidJWTFormat,
			wantErr:     true,
		},
		{
			name:     "Invalid base64 in header",
			jwtToken: "!!!.eyJzdWIiOiIxMjM0NTY3ODkwIn0.sig",
			wantErr:  true,
		},
		{
			name:     "Invalid base64 in payload",
			jwtToken: "eyJhbGciOiJIUzI1NiJ9.!!!.sig",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, payload, err := DecodeRaw(tt.jwtToken)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.ErrorIs(t, err, tt.expectedErr)
				}
				return
			}

			assert.NoError(t, err)
			assert.JSONEq(t, string(tt.expectedHeader), string(header))
			assert.JSONEq(t, string(tt.expectedBody), string(payload))
		})
	}
}

func Test_Decode(t *testing.T) {
	tests := []struct {
		name            string
		jwtToken        string
		expectedHeader  map[string]any
		expectedPayload map[string]any
		expectedErr     error
		wantErr         bool
	}{
		{
			name:     "Valid JWT token",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30",
			expectedHeader: map[string]any{
				"alg": "HS256",
				"typ": "JWT",
			},
			expectedPayload: map[string]any{
				"name":  "John Doe",
				"sub":   "1234567890",
				"admin": true,
				"iat":   float64(1516239022),
			},
			wantErr: false,
		},
		{
			name:        "Invalid JWT format",
			jwtToken:    "abc123",
			expectedErr: ErrInvalidJWTFormat,
			wantErr:     true,
		},
		{
			name:     "Invalid base64 in header",
			jwtToken: "!!!.eyJzdWIiOiIxMjM0NTY3ODkwIn0.sig",
			wantErr:  true,
		},
		{
			name:     "Invalid JSON in payload",
			jwtToken: "eyJhbGciOiJIUzI1NiJ9.bm90anNvbg.sig",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, payload, err := Decode(tt.jwtToken)

			if tt.wantErr {
				assert.Error(t, err)

				if tt.expectedErr != nil {
					assert.ErrorIs(t, err, tt.expectedErr)
				}

				assert.Nil(t, header)
				assert.Nil(t, payload)
				return
			}

			assert.NoError(t, err)

			// Strong equality check (catches missing/extra keys)
			assert.Equal(t, tt.expectedHeader, header)
			assert.Equal(t, tt.expectedPayload, payload)
		})
	}
}
