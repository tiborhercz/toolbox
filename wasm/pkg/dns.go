package pkg

import (
	"context"
	"encoding/json"
	"syscall/js"

	"github.com/tiborhercz/toolbox/pkg/dns"
)

func DnsLookup(this js.Value, args []js.Value) interface{} {
	domain := args[0].String()
	recordType := args[1].String()

	handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) interface{} {
		resolve := promiseArgs[0]
		reject := promiseArgs[1]

		go func() {
			r := dns.NewDefaultResolver()
			ctx := context.Background()

			var records []dns.Record
			var err error

			switch recordType {
			case "", "ALL":
				records, err = r.LookupAll(ctx, domain)
			case "DMARC":
				records, err = r.LookupDMARC(ctx, domain)
			default:
				records, err = r.Lookup(ctx, domain, recordType)
			}

			if err != nil {
				reject.Invoke(err.Error())
				return
			}

			jsonData, err := json.Marshal(records)
			if err != nil {
				reject.Invoke(err.Error())
				return
			}

			resolve.Invoke(string(jsonData))
		}()

		return nil
	})

	return js.Global().Get("Promise").New(handler)
}

func DnsPropagation(this js.Value, args []js.Value) interface{} {
	domain := args[0].String()
	recordType := args[1].String()

	handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) interface{} {
		resolve := promiseArgs[0]
		reject := promiseArgs[1]

		go func() {
			ctx := context.Background()
			results := dns.CheckPropagation(ctx, domain, recordType)
			summary := dns.ComparePropagation(results)

			response := struct {
				Results []dns.PropagationResult `json:"results"`
				Summary dns.PropagationSummary  `json:"summary"`
			}{
				Results: results,
				Summary: summary,
			}

			jsonData, err := json.Marshal(response)
			if err != nil {
				reject.Invoke(err.Error())
				return
			}

			resolve.Invoke(string(jsonData))
		}()

		return nil
	})

	return js.Global().Get("Promise").New(handler)
}
