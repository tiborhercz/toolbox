package dns

import (
	"context"
	"strings"
	"sync"
)

var defaultRecordTypes = []string{"A", "AAAA", "MX", "NS", "TXT", "CNAME"}

type namedProvider struct {
	name     string
	provider Provider
}

var defaultPropagationProviders = []namedProvider{
	{"Cloudflare", NewCloudflareProvider()},
	{"Google", NewGoogleProvider()},
	{"Quad9", NewQuad9Provider()},
}

type PropagationResult struct {
	Provider string   `json:"provider"`
	Records  []Record `json:"records"`
	Error    string   `json:"error,omitempty"`
}

type Resolver struct {
	provider Provider
}

func NewResolver(provider Provider) *Resolver {
	return &Resolver{provider: provider}
}

func NewDefaultResolver() *Resolver {
	return NewResolver(NewCloudflareProvider())
}

func (r *Resolver) Lookup(ctx context.Context, domain, recordType string) ([]Record, error) {
	return r.provider.Query(ctx, domain, recordType)
}

func (r *Resolver) LookupDMARC(ctx context.Context, domain string) ([]Record, error) {
	records, err := r.provider.Query(ctx, "_dmarc."+domain, "TXT")
	if err != nil {
		return nil, err
	}

	for i := range records {
		records[i].Type = "DMARC"
	}
	return records, nil
}

func (r *Resolver) LookupAll(ctx context.Context, domain string) ([]Record, error) {
	var all []Record

	for _, rt := range defaultRecordTypes {
		records, err := r.provider.Query(ctx, domain, rt)
		if err != nil {
			continue
		}
		all = append(all, records...)
	}

	if dmarc, err := r.LookupDMARC(ctx, domain); err == nil {
		all = append(all, dmarc...)
	}

	return all, nil
}

func CheckPropagation(ctx context.Context, domain, recordType string) []PropagationResult {
	results := make([]PropagationResult, len(defaultPropagationProviders))
	var wg sync.WaitGroup

	for i, np := range defaultPropagationProviders {
		wg.Add(1)
		go func(idx int, np namedProvider) {
			defer wg.Done()

			r := NewResolver(np.provider)
			var records []Record
			var err error

			switch recordType {
			case "", "ALL":
				records, err = r.LookupAll(ctx, domain)
			case "DMARC":
				records, err = r.LookupDMARC(ctx, domain)
			default:
				records, err = r.Lookup(ctx, domain, recordType)
			}

			result := PropagationResult{Provider: np.name}
			if err != nil {
				result.Error = err.Error()
			} else {
				result.Records = records
			}
			results[idx] = result
		}(i, np)
	}

	wg.Wait()
	return results
}

func ComparePropagation(results []PropagationResult) PropagationSummary {
	dataByProvider := make(map[string]map[string]bool)

	for _, r := range results {
		if r.Error != "" {
			continue
		}
		values := make(map[string]bool)
		for _, rec := range r.Records {
			values[normalizeRecordKey(rec)] = true
		}
		dataByProvider[r.Provider] = values
	}

	allKeys := make(map[string]bool)
	for _, values := range dataByProvider {
		for k := range values {
			allKeys[k] = true
		}
	}

	consistent := true
	var diffs []PropagationDiff

	for key := range allKeys {
		var have, missing []string
		for _, r := range results {
			if r.Error != "" {
				continue
			}
			if dataByProvider[r.Provider][key] {
				have = append(have, r.Provider)
			} else {
				missing = append(missing, r.Provider)
			}
		}
		if len(missing) > 0 {
			consistent = false
			diffs = append(diffs, PropagationDiff{
				Record:  key,
				Present: have,
				Missing: missing,
			})
		}
	}

	return PropagationSummary{
		Consistent: consistent,
		Diffs:      diffs,
	}
}

// normalizeRecordKey builds a comparison key with quoting differences stripped
// so providers that wrap TXT/DMARC values in quotes match those that don't.
func normalizeRecordKey(r Record) string {
	data := strings.Trim(r.Data, "\"")
	return r.Type + ":" + data
}

type PropagationSummary struct {
	Consistent bool              `json:"consistent"`
	Diffs      []PropagationDiff `json:"diffs,omitempty"`
}

type PropagationDiff struct {
	Record  string   `json:"record"`
	Present []string `json:"present"`
	Missing []string `json:"missing"`
}
