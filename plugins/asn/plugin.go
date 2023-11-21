package asn

import (
	"fmt"

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
	iptec.Cash
}

func (p *Asn) Name() string {
	return p.name
}

func (p *Asn) Activate() error {
	p.Log.Info("ACTIVATION " + p.Name())
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
	// download()
	return nil
}

func (p *Asn) Find(address string) (interface{}, error) {
	a := AsnResult{
		asn: "test",
	}

	return a, nil
}
