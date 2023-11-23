package dnsbl

import (
	"fmt"
	"net"
	"regexp"
	"time"

	"github.com/miekg/dns"
)

type addressResult struct {
	dnsbl  string
	exist  bool
	reason string
}

type dnsblCLient struct {
	client      *dns.Client
	resolver    string
	answerRegex *regexp.Regexp
}

func newDnsblClient(resolver string) (*dnsblCLient, error) {

	client := new(dns.Client)

	client.DialTimeout = (2 * time.Second)
	client.ReadTimeout = (5 * time.Second)
	client.WriteTimeout = (2 * time.Second)

	return &dnsblCLient{
		client:      client,
		resolver:    fmt.Sprintf("%s:53", resolver),
		answerRegex: regexp.MustCompile(`regexp`),
	}, nil
}

func (d *dnsblCLient) check(query string) bool {
	msg := new(dns.Msg)
	msg.SetQuestion(query, dns.TypeA)

	response, _, err := d.client.Exchange(msg, d.resolver)
	if err != nil {
		return false
	}
	for _, ans := range response.Answer {
		if record, ok := ans.(*dns.A); ok {
			answerAddress := net.ParseIP(record.A.String())
			if answerAddress != nil {
				segments := answerAddress.To4()
				if segments != nil && segments[3] > 1 {
					return true
				}
			}
		}
	}
	return false
}

func (d *dnsblCLient) reason(query string) string {
	msg := new(dns.Msg)
	msg.SetQuestion(query, dns.TypeTXT)
	response, _, err := d.client.Exchange(msg, d.resolver)
	if err != nil {
		return ""
	}
	for _, ans := range response.Answer {
		if t, ok := ans.(*dns.TXT); ok {
			return t.Txt[0]
		}
	}
	return ""
}

func (d *dnsblCLient) Find(reversedAddress, provider string) *addressResult {
	query := fmt.Sprintf("%s.%s.", reversedAddress, provider)
	if d.check(query) {
		return &addressResult{
			exist:  true,
			reason: d.reason(query),
		}
	}
	return nil
}
