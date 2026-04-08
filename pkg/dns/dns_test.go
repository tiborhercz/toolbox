package dns

import (
	"context"
	"testing"
)

const testDomain = "example.com"

func newTestResolver() *Resolver {
	return NewDefaultResolver()
}

func TestLookupA(t *testing.T) {
	r := newTestResolver()
	records, err := r.Lookup(context.Background(), testDomain, "A")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one A record for example.com")
	}

	for _, rec := range records {
		if rec.Type != "A" {
			t.Fatalf("expected record type A, got %s", rec.Type)
		}
	}
}

func TestLookupAAAA(t *testing.T) {
	r := newTestResolver()
	records, err := r.Lookup(context.Background(), testDomain, "AAAA")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one AAAA record for example.com")
	}

	for _, rec := range records {
		if rec.Type != "AAAA" {
			t.Fatalf("expected record type AAAA, got %s", rec.Type)
		}
	}
}

func TestLookupMX(t *testing.T) {
	r := newTestResolver()
	records, err := r.Lookup(context.Background(), "google.com", "MX")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one MX record for google.com")
	}
}

func TestLookupNS(t *testing.T) {
	r := newTestResolver()
	records, err := r.Lookup(context.Background(), testDomain, "NS")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one NS record for example.com")
	}
}

func TestLookupTXT(t *testing.T) {
	r := newTestResolver()
	records, err := r.Lookup(context.Background(), "google.com", "TXT")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one TXT record for google.com")
	}
}

func TestLookupCNAME(t *testing.T) {
	r := newTestResolver()
	_, err := r.Lookup(context.Background(), testDomain, "CNAME")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}
}

func TestLookupAll(t *testing.T) {
	r := newTestResolver()
	records, err := r.LookupAll(context.Background(), testDomain)
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one record for example.com")
	}

	typesSeen := make(map[string]bool)
	for _, rec := range records {
		typesSeen[rec.Type] = true
	}

	if !typesSeen["A"] && !typesSeen["AAAA"] {
		t.Fatal("expected at least one A or AAAA record in LookupAll results")
	}
}

func TestProviderInterface(t *testing.T) {
	var p Provider = NewCloudflareProvider()
	records, err := p.Query(context.Background(), testDomain, "A")
	if err != nil {
		t.Skip("DNS lookup failed, skipping due to possible network issue:", err)
	}

	if len(records) == 0 {
		t.Fatal("expected at least one A record via Provider interface")
	}
}
