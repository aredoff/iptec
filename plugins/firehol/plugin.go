package firehol

import (
	"fmt"
	"net"
	"strings"

	"github.com/aredoff/iptec"
)

const (
	name = "firehol"
)

func New() *Firehol {
	// log.Error("Blocking request. Unable to parse source address")
	return &Firehol{
		name: name,
	}
}

type Firehol struct {
	name string
	iptec.Сurator
	iptec.Cash
	iptec.WebClient
	data *fireholData
}

func (p *Firehol) Name() string {
	return p.name
}

func (p *Firehol) Activate() error {
	p.data = newFireholData()

	data, err := p.Cash.Get("firehol")
	if err == nil {
		p.Log.Info("Load data from cash.")
		p.data.Deserialization(data)
		return nil
	}
	p.Log.Info("Collect data from firehol.org.")
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

			} else {
				_, network, err := net.ParseCIDR(line)
				if err != nil {
					continue
				}
				p.data.AddNet(name, *network)
			}
		}
	}

	data, err = p.data.Serialization()
	if err != nil {
		return err
	}
	err = p.Cash.Set("firehol", data)
	if err != nil {
		return fmt.Errorf("cant set in cash data, err=%s", err)
	}
	return nil
}

func (p *Firehol) Find(address net.IP) (iptec.PluginReport, error) {
	list := p.data.Find(address)

	a := fireholResult{
		Name:  "firehol",
		Lists: list,
	}
	return &a, nil
}
