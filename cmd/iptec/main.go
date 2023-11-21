package main

import (
	"fmt"

	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/asn"
	"github.com/aredoff/iptec/plugins/dns"
)

func main() {
	a := iptec.New()
	defer a.Close()
	a.Use(asn.New())
	a.Use(dns.New())
	a.Activate()
	report, err := a.Find("8.8.8.8")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(report)
}
