package dnsbl

import (
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/aredoff/iptec"
)

const (
	name = "dnsbl"
)

type providerResult struct {
	dnsbl  string
	exist  bool
	reason string
}

func New() *Dnsbl {
	return &Dnsbl{
		name: name,
	}
}

type Dnsbl struct {
	name string
	iptec.DnsClient
}

func (p *Dnsbl) Name() string {
	return p.name
}

func (p *Dnsbl) Activate() error {
	return nil
}

func (p *Dnsbl) Find(address net.IP) (iptec.PluginReport, error) {
	reverseAddress, err := reverseIP(address)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	ch := make(chan *providerResult, 16)
	var list []string
	if address.To4() != nil {
		list = sourcesIpv4
	} else {
		list = sourcesIpv6
	}
	for _, v := range list {
		wg.Add(1)
		go func(provider string) {
			defer wg.Done()
			var res *providerResult
			recordsA, err := p.Dns.A(fmt.Sprintf("%s.%s.", reverseAddress, provider))
			if err == nil && len(recordsA) > 0 {
				if checkArecord(recordsA[0]) {
					res = &providerResult{
						dnsbl:  provider,
						exist:  true,
						reason: "",
					}
					recoredsTXT, err := p.Dns.TXT(fmt.Sprintf("%s.%s.", reverseAddress, provider))
					if err == nil && len(recoredsTXT) > 0 {
						res.reason = strings.Join(recoredsTXT, ",")
					}
				}
			}
			if res != nil {
				ch <- res
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	report := &dnsblResult{
		Lists: map[string]string{},
	}

	for result := range ch {
		report.Lists[result.dnsbl] = result.reason
	}
	return report, nil
}
