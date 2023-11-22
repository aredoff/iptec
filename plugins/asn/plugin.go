package asn

import (
	"fmt"
	"net"

	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/log"
)

func New() *Asn {
	return &Asn{
		name: "asn",
	}
}

type Asn struct {
	name string
	iptec.Сurator
	iptec.Cash
	iptec.WebClient
}

func (p *Asn) Name() string {
	return p.name
}

func (p *Asn) Activate() error {
	p.Log.Info("Activation: ok")
	// res, err := net.LookupTXT("8.8.8.8")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := p.Cash.Set("test", []byte("test OK!!!!!!"))
	if err != nil {
		return fmt.Errorf("cant set cash err=%s", err)
	}
	res2, err := p.Cash.Get("test")
	if err != nil {
		return fmt.Errorf("cant get cash err=%s", err)
	}
	log.Warning(string(res2))
	p.Client.Get("https://google.ru")
	// download()
	return nil
}

func (p *Asn) Find(address net.IP) (AsnResult, error) {
	a := AsnResult{
		asn: "test",
	}

	return a, nil
}
