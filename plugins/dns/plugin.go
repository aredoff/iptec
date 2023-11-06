package dns

import (
	"github.com/aredoff/iptec"
	clog "github.com/aredoff/iptec/log"
)

var log = clog.NewWithPlugin("dns")

func New() *Dns {
	// log.Error("Blocking request. Unable to parse source address")
	return &Dns{
		name: "dns",
	}
}

type Dns struct {
	name string
	iptec.СuratorMixin
	iptec.CashMixin
}

func (p *Dns) Name() string {
	return p.name
}

func (p *Dns) Activate() {
	log.Warning("ACTIVATION " + p.Name())
}
