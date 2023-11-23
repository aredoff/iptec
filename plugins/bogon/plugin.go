package bogon

import (
	"fmt"
	"net"

	"github.com/aredoff/iptec"
)

const (
	name = "bogon"
)

func New() *Bogon {
	return &Bogon{
		name:   name,
		bogons: map[string]net.IPNet{},
	}
}

type Bogon struct {
	name   string
	bogons map[string]net.IPNet
}

func (p *Bogon) Name() string {
	return p.name
}

func (p *Bogon) Activate() error {
	for _, v := range nets {
		_, network, err := net.ParseCIDR(v)
		if err != nil {
			return fmt.Errorf("cant parse bogon %s, err=%s", v, err)
		}
		if v != network.String() {
			fmt.Println(v)
		}
		p.bogons[network.String()] = *network
	}
	return nil
}

func (p *Bogon) Find(address net.IP) (iptec.PluginReport, error) {
	isBogon := false
	for _, v := range p.bogons {
		if v.Contains(address) {
			isBogon = true
		}
	}
	return &bogonResult{
		Bogon: isBogon,
	}, nil
}
