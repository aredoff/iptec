package tor

import (
	"fmt"
	"net"
	"strings"

	"github.com/aredoff/iptec"
)

const (
	name  = "tor"
	dnsel = "dnsel.torproject.org"
)

func New() *Tor {
	return &Tor{
		name: name,
	}
}

type Tor struct {
	name string
	data *torData
	iptec.Сurator
	iptec.WebClient
	iptec.DnsClient
	iptec.Cash
}

func (p *Tor) Name() string {
	return p.name
}

func (p *Tor) Activate() error {
	p.data = newTorData()

	data, err := p.Cash.Get("tor")
	if err == nil {
		p.Log.Info("Load data from cash.")
		err := p.data.Deserialization(data)
		if err != nil {
			return err
		}
		return nil
	}
	p.Log.Info("Collect data from sources")
	for name, url := range sources {
		r, err := p.Client.Get(url)
		if err != nil {
			return fmt.Errorf("cant load %s, err=%s", name, err)
		}
		if r.StatusCode != 200 {
			return fmt.Errorf("cant load %s, status_code=%d", name, r.StatusCode)
		}
		for _, line := range strings.Split(string(r.Body), "\n") {
			address := net.ParseIP(line)
			if address != nil {
				p.data.AddIp(name, address)

			}
		}
	}

	data, err = p.data.Serialization()
	if err != nil {
		return err
	}
	err = p.Cash.Set("tor", data)
	if err != nil {
		return fmt.Errorf("cant set in cash data, err=%s", err)
	}
	return nil
}

func (p *Tor) Find(address net.IP) (iptec.PluginReport, error) {
	report := &torResult{}

	reverseAddress, err := reverseIP(address)
	if err != nil {
		return nil, err
	}
	recordsA, err := p.Dns.A(fmt.Sprintf("%s.%s.", reverseAddress, dnsel))
	if err == nil && len(recordsA) > 0 {
		if checkArecord(recordsA[0]) {
			report.Detect = true
			report.Proofs = []string{dnsel}
			return report, nil
		}
	}

	finds := p.data.Find(address)
	if len(finds) > 0 {
		report.Detect = true
		report.Proofs = finds
	}

	return report, nil
}
