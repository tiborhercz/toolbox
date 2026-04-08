package dns

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/dns/dnsmessage"
)

const dnsTypeCAA = dnsmessage.Type(257)

var wireTypeMap = map[string]dnsmessage.Type{
	"A":     dnsmessage.TypeA,
	"AAAA":  dnsmessage.TypeAAAA,
	"CNAME": dnsmessage.TypeCNAME,
	"MX":    dnsmessage.TypeMX,
	"NS":    dnsmessage.TypeNS,
	"SOA":   dnsmessage.TypeSOA,
	"TXT":   dnsmessage.TypeTXT,
	"SRV":   dnsmessage.TypeSRV,
	"CAA":   dnsTypeCAA,
}

type Quad9Provider struct {
	client  *http.Client
	baseURL string
}

func NewQuad9Provider() *Quad9Provider {
	return &Quad9Provider{
		client:  &http.Client{Timeout: 5 * time.Second},
		baseURL: "https://dns.quad9.net/dns-query",
	}
}

func (p *Quad9Provider) Name() string {
	return "Quad9"
}

func (p *Quad9Provider) Query(ctx context.Context, domain, recordType string) ([]Record, error) {
	qtype, ok := wireTypeMap[recordType]
	if !ok {
		return nil, fmt.Errorf("unsupported record type for wire format: %s", recordType)
	}

	name, err := dnsmessage.NewName(domain + ".")
	if err != nil {
		return nil, fmt.Errorf("invalid domain: %w", err)
	}

	msg := dnsmessage.Message{
		Header: dnsmessage.Header{RecursionDesired: true},
		Questions: []dnsmessage.Question{{
			Name:  name,
			Type:  qtype,
			Class: dnsmessage.ClassINET,
		}},
	}

	wireQuery, err := msg.Pack()
	if err != nil {
		return nil, fmt.Errorf("packing DNS query: %w", err)
	}

	encoded := base64.RawURLEncoding.EncodeToString(wireQuery)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.baseURL+"?dns="+encoded, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Accept", "application/dns-message")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("querying DoH (Quad9): %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response (Quad9): %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DoH (Quad9) returned status %d", resp.StatusCode)
	}

	return parseWireResponse(body)
}

func parseWireResponse(data []byte) ([]Record, error) {
	var parser dnsmessage.Parser
	header, err := parser.Start(data)
	if err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	if header.RCode != dnsmessage.RCodeSuccess {
		return nil, fmt.Errorf("DNS query failed with RCODE %d", header.RCode)
	}

	if err := parser.SkipAllQuestions(); err != nil {
		return nil, fmt.Errorf("skipping questions: %w", err)
	}

	var records []Record
	for {
		h, err := parser.AnswerHeader()
		if err == dnsmessage.ErrSectionDone {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading answer header: %w", err)
		}

		rec := Record{
			Name: strings.TrimSuffix(h.Name.String(), "."),
			Type: dnsTypeName(int(h.Type)),
			TTL:  h.TTL,
		}

		data, err := parseResourceData(&parser, h.Type)
		if err != nil {
			return nil, err
		}
		if data == "" {
			continue
		}

		rec.Data = data
		records = append(records, rec)
	}

	return records, nil
}

func parseResourceData(parser *dnsmessage.Parser, t dnsmessage.Type) (string, error) {
	switch t {
	case dnsmessage.TypeA:
		r, err := parser.AResource()
		if err != nil {
			return "", err
		}
		return net.IP(r.A[:]).String(), nil
	case dnsmessage.TypeAAAA:
		r, err := parser.AAAAResource()
		if err != nil {
			return "", err
		}
		return net.IP(r.AAAA[:]).String(), nil
	case dnsmessage.TypeCNAME:
		r, err := parser.CNAMEResource()
		if err != nil {
			return "", err
		}
		return r.CNAME.String(), nil
	case dnsmessage.TypeMX:
		r, err := parser.MXResource()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d %s", r.Pref, r.MX.String()), nil
	case dnsmessage.TypeNS:
		r, err := parser.NSResource()
		if err != nil {
			return "", err
		}
		return r.NS.String(), nil
	case dnsmessage.TypeTXT:
		r, err := parser.TXTResource()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("\"%s\"", strings.Join(r.TXT, "")), nil
	case dnsmessage.TypeSOA:
		r, err := parser.SOAResource()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s %d %d %d %d %d",
			r.NS.String(), r.MBox.String(), r.Serial, r.Refresh, r.Retry, r.Expire, r.MinTTL), nil
	case dnsmessage.TypeSRV:
		r, err := parser.SRVResource()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d %d %d %s", r.Priority, r.Weight, r.Port, r.Target.String()), nil
	case dnsTypeCAA:
		r, err := parser.UnknownResource()
		if err != nil {
			return "", err
		}
		return parseCAAData(r.Data), nil
	default:
		return "", nil
	}
}

// CAA wire format: [1 byte flags] [1 byte tag length] [tag] [value]
func parseCAAData(data []byte) string {
	if len(data) < 2 {
		return ""
	}
	flags := data[0]
	tagLen := int(data[1])
	if len(data) < 2+tagLen {
		return ""
	}
	tag := string(data[2 : 2+tagLen])
	value := string(data[2+tagLen:])
	return fmt.Sprintf("%d %s \"%s\"", flags, tag, value)
}
