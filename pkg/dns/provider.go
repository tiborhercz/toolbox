package dns

import "context"

type Provider interface {
	Query(ctx context.Context, domain string, recordType string) ([]Record, error)
}

type Record struct {
	Name string `json:"name"`
	Type string `json:"type"`
	TTL  uint32 `json:"ttl"`
	Data string `json:"data"`
}
