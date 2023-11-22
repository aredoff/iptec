package dns

import (
	"net"

	"github.com/aredoff/iptec"
)

func New() *Dns {
	// log.Error("Blocking request. Unable to parse source address")
	return &Dns{
		name: "dns",
	}
}

type Dns struct {
	name string
	iptec.Сurator
	iptec.Cash
}

func (p *Dns) Name() string {
	return p.name
}

func (p *Dns) Activate() error {
	p.Log.Info("Activation: ok")
	return nil
}

func (p *Dns) Find(address net.IP) (DnsResult, error) {
	a := DnsResult{
		asn: "test",
	}
	return a, nil
}
