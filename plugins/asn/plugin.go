package asn

import (
	"github.com/aredoff/iptec"
	clog "github.com/aredoff/iptec/log"
)

var log = clog.NewWithPlugin("asn")

func New() *Asn {
	return &Asn{
		name: "asn",
	}
}

type Asn struct {
	name string
	iptec.СuratorMixin
	iptec.CashMixin
}

func (p *Asn) Name() string {
	return p.name
}

func (p *Asn) Activate() {
	log.Warning("1 " + p.Name())
	log.Warning("ACTIVATION " + p.Name())
	// res, err := net.LookupTXT("8.8.8.8")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := p.CashSet("test", []byte("test OK!!!!!!"))
	if err != nil {
		log.Warning("cant set")
	}
	res2, err := p.CashGet("test")
	if err != nil {
		log.Warning("cant get")
	}
	log.Warning(string(res2))
	// download()
}
