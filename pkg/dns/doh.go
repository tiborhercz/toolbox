package dns

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// IANA DNS Resource Record type codes: https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4
var dnsTypeNames = map[int]string{
	1:   "A",
	2:   "NS",
	5:   "CNAME",
	6:   "SOA",
	15:  "MX",
	16:  "TXT",
	28:  "AAAA",
	33:  "SRV",
	257: "CAA",
}

type DoHProvider struct {
	name    string
	client  *http.Client
	baseURL string
}

type dohResponse struct {
	Status int         `json:"Status"`
	Answer []dohAnswer `json:"Answer"`
}

type dohAnswer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  uint32 `json:"TTL"`
	Data string `json:"data"`
}

func NewCloudflareProvider() *DoHProvider {
	return &DoHProvider{
		name:    "Cloudflare",
		client:  &http.Client{Timeout: 5 * time.Second},
		baseURL: "https://cloudflare-dns.com/dns-query",
	}
}

func NewGoogleProvider() *DoHProvider {
	return &DoHProvider{
		name:    "Google",
		client:  &http.Client{Timeout: 5 * time.Second},
		baseURL: "https://dns.google/resolve",
	}
}

func (p *DoHProvider) Name() string {
	return p.name
}

func (p *DoHProvider) Query(ctx context.Context, domain, recordType string) ([]Record, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	q := req.URL.Query()
	q.Set("name", domain)
	q.Set("type", recordType)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Accept", "application/dns-json")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("querying DoH (%s): %w", p.name, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response (%s): %w", p.name, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DoH (%s) returned status %d: %s", p.name, resp.StatusCode, body)
	}

	var dohResp dohResponse
	if err := json.Unmarshal(body, &dohResp); err != nil {
		return nil, fmt.Errorf("parsing response (%s): %w", p.name, err)
	}

	if dohResp.Status != 0 {
		return nil, fmt.Errorf("DNS query failed (%s) with status %d (RCODE)", p.name, dohResp.Status)
	}

	records := make([]Record, 0, len(dohResp.Answer))
	for _, a := range dohResp.Answer {
		records = append(records, Record{
			Name: a.Name,
			Type: dnsTypeName(a.Type),
			TTL:  a.TTL,
			Data: a.Data,
		})
	}

	return records, nil
}

func dnsTypeName(typeCode int) string {
	if name, ok := dnsTypeNames[typeCode]; ok {
		return name
	}
	return strconv.Itoa(typeCode)
}
