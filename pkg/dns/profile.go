package dns

import (
	"context"
	"sync"
	"time"
)

type ProfileResult struct {
	Provider string        `json:"provider"`
	Records  []Record      `json:"records"`
	Latency  time.Duration `json:"latency_ns"`
	Error    string        `json:"error,omitempty"`
}

type ProfileProvider struct {
	Name     string
	Provider Provider
	// InitErr is set when constructing the provider failed (e.g. system DNS
	// could not be detected); the Profile run will record this as the error
	// for the result without attempting a query.
	InitErr error
}

// DefaultProfileProviders returns the public DoH providers plus the host's
// system resolver. If the system resolver cannot be detected the entry is
// still returned with InitErr set so the user can see why it was skipped.
func DefaultProfileProviders() []ProfileProvider {
	providers := []ProfileProvider{
		{Name: "Cloudflare", Provider: NewCloudflareProvider()},
		{Name: "Google", Provider: NewGoogleProvider()},
		{Name: "Quad9", Provider: NewQuad9Provider()},
		{Name: "NextDNS", Provider: NewNextDNSProvider()},
		{Name: "Control D", Provider: NewControlDProvider()},
	}

	sys, err := NewSystemProvider()
	if err != nil {
		providers = append(providers, ProfileProvider{Name: "System", InitErr: err})
	} else {
		providers = append(providers, ProfileProvider{Name: sys.Name(), Provider: sys})
	}
	return providers
}

// Profile runs the given query against every provider concurrently and
// measures the per-provider latency.
func Profile(ctx context.Context, domain, recordType string, providers []ProfileProvider) []ProfileResult {
	results := make([]ProfileResult, len(providers))
	var wg sync.WaitGroup

	for i, p := range providers {
		wg.Add(1)
		go func(idx int, pp ProfileProvider) {
			defer wg.Done()

			result := ProfileResult{Provider: pp.Name}
			if pp.InitErr != nil {
				result.Error = pp.InitErr.Error()
				results[idx] = result
				return
			}

			start := time.Now()
			records, err := pp.Provider.Query(ctx, domain, recordType)
			result.Latency = time.Since(start)

			if err != nil {
				result.Error = err.Error()
			} else {
				result.Records = records
			}
			results[idx] = result
		}(i, p)
	}

	wg.Wait()
	return results
}
