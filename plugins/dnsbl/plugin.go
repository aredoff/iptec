package dnsbl

import (
	"net"
	"sync"

	"github.com/aredoff/iptec"
)

const (
	name = "dnsbl"
)

func New(resolver string) *Dnsbl {
	return &Dnsbl{
		name:     name,
		resolver: resolver,
	}
}

type Dnsbl struct {
	name     string
	resolver string
	client   *dnsblCLient
}

func (p *Dnsbl) Name() string {
	return p.name
}

func (p *Dnsbl) Activate() error {
	var err error
	p.client, err = newDnsblClient(p.resolver)
	if err != nil {
		return err
	}
	return nil
}

func (p *Dnsbl) Find(address net.IP) (iptec.PluginReport, error) {
	reverseAddress, err := reverseIP(address)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	ch := make(chan *addressResult, 16)
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
			rep := p.client.Find(reverseAddress, provider)
			if rep != nil {
				rep.dnsbl = provider
				ch <- rep
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
