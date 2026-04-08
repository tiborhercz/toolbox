package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/toolbox/internal/model"
	"github.com/tiborhercz/toolbox/pkg/dns"
)

var (
	dnsOptions model.DnsOptions

	dnsCmd = &cobra.Command{
		Use:   "dns",
		Short: "DNS lookup tools",
	}

	dnsLookupCmd = &cobra.Command{
		Use:   "lookup [domain]",
		Short: "Look up DNS records for a domain",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Fatal("requires a domain argument. Example: toolbox dns lookup example.com")
			}

			domain := args[0]
			ctx := context.Background()
			r := dns.NewDefaultResolver()

			recordType := strings.ToUpper(dnsOptions.RecordType)

			var records []dns.Record
			var err error

			switch recordType {
			case "":
				records, err = r.LookupAll(ctx, domain)
			case "DMARC":
				records, err = r.LookupDMARC(ctx, domain)
			default:
				records, err = r.Lookup(ctx, domain, recordType)
			}

			if err != nil {
				logrus.Fatal(err)
			}

			if len(records) == 0 {
				logrus.Info("No records found")
				return
			}

			logrus.Info(formatRecords(records))
		},
	}
	dnsPropagationCmd = &cobra.Command{
		Use:   "propagation [domain]",
		Short: "Check DNS propagation across multiple providers (Cloudflare, Google, Quad9)",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Fatal("requires a domain argument. Example: toolbox dns propagation example.com")
			}

			domain := args[0]
			ctx := context.Background()
			recordType := strings.ToUpper(dnsOptions.RecordType)

			results := dns.CheckPropagation(ctx, domain, recordType)

			for i, result := range results {
				if i > 0 {
					fmt.Println()
				}
				fmt.Printf("=== %s ===\n", result.Provider)
				if result.Error != "" {
					logrus.Errorf("  Error: %s", result.Error)
					continue
				}
				if len(result.Records) == 0 {
					fmt.Println("  No records found")
					continue
				}
				fmt.Println(formatRecords(result.Records))
			}

			summary := dns.ComparePropagation(results)
			fmt.Println()
			if summary.Consistent {
				fmt.Println("=== Result: All providers returned consistent records ===")
			} else {
				fmt.Println("=== Result: Differences detected ===")
				for _, d := range summary.Diffs {
					fmt.Printf("  %s\n    Present: %s\n    Missing: %s\n",
						d.Record, strings.Join(d.Present, ", "), strings.Join(d.Missing, ", "))
				}
			}
		},
	}
)

var supportedRecordTypes = []string{"A", "AAAA", "MX", "NS", "TXT", "CNAME", "SOA", "CAA", "DMARC"}

func formatRecords(records []dns.Record) string {
	grouped := make(map[string][]dns.Record)
	var order []string

	for _, r := range records {
		if _, exists := grouped[r.Type]; !exists {
			order = append(order, r.Type)
		}
		grouped[r.Type] = append(grouped[r.Type], r)
	}

	var b strings.Builder
	for i, t := range order {
		if i > 0 {
			b.WriteString("\n")
		}
		fmt.Fprintf(&b, "%s Records:\n", t)
		for _, r := range grouped[t] {
			fmt.Fprintf(&b, "  %s (TTL: %d)\n", r.Data, r.TTL)
		}
	}

	return strings.TrimRight(b.String(), "\n")
}

func init() {
	rootCmd.AddCommand(dnsCmd)
	dnsCmd.AddCommand(dnsLookupCmd, dnsPropagationCmd)

	typeFlag := "DNS record type to query. Supported types: " + strings.Join(supportedRecordTypes, ", ")
	dnsLookupCmd.Flags().StringVarP(&dnsOptions.RecordType, "type", "t", "", typeFlag)
	dnsPropagationCmd.Flags().StringVarP(&dnsOptions.RecordType, "type", "t", "", typeFlag)
}
