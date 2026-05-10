package dns

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/net/dns/dnsmessage"
)

const resolvConfPath = "/etc/resolv.conf"

type SystemProvider struct {
	server  string
	timeout time.Duration
}

// NewSystemProvider returns a provider that queries the host's configured DNS
// server (typically the one issued by DHCP). On Unix-like systems it reads
// /etc/resolv.conf; on platforms without that file it returns an error.
func NewSystemProvider() (*SystemProvider, error) {
	server, err := detectSystemDNSServer()
	if err != nil {
		return nil, err
	}
	return &SystemProvider{server: server, timeout: 5 * time.Second}, nil
}

func (p *SystemProvider) Name() string {
	return fmt.Sprintf("System (%s)", p.server)
}

func (p *SystemProvider) Server() string {
	return p.server
}

func (p *SystemProvider) Query(ctx context.Context, domain, recordType string) ([]Record, error) {
	qtype, ok := wireTypeMap[recordType]
	if !ok {
		return nil, fmt.Errorf("unsupported record type for system resolver: %s", recordType)
	}

	name, err := dnsmessage.NewName(domain + ".")
	if err != nil {
		return nil, fmt.Errorf("invalid domain: %w", err)
	}

	msg := dnsmessage.Message{
		Header: dnsmessage.Header{
			ID:               uint16(time.Now().UnixNano()),
			RecursionDesired: true,
		},
		Questions: []dnsmessage.Question{{
			Name:  name,
			Type:  qtype,
			Class: dnsmessage.ClassINET,
		}},
	}

	wire, err := msg.Pack()
	if err != nil {
		return nil, fmt.Errorf("packing DNS query: %w", err)
	}

	d := net.Dialer{Timeout: p.timeout}
	conn, err := d.DialContext(ctx, "udp", net.JoinHostPort(p.server, "53"))
	if err != nil {
		return nil, fmt.Errorf("dialing %s: %w", p.server, err)
	}
	defer conn.Close()

	deadline := time.Now().Add(p.timeout)
	if dl, ok := ctx.Deadline(); ok && dl.Before(deadline) {
		deadline = dl
	}
	_ = conn.SetDeadline(deadline)

	if _, err := conn.Write(wire); err != nil {
		return nil, fmt.Errorf("writing query: %w", err)
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	return parseWireResponse(buf[:n])
}

func detectSystemDNSServer() (string, error) {
	f, err := os.Open(resolvConfPath)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", resolvConfPath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "nameserver" {
			return fields[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanning %s: %w", resolvConfPath, err)
	}
	return "", fmt.Errorf("no nameserver entry found in %s", resolvConfPath)
}
