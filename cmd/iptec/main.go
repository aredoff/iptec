package main

import (
	"encoding/json"
	"fmt"

	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/bogon"
	"github.com/aredoff/iptec/plugins/dnsbl"
	"github.com/aredoff/iptec/plugins/firehol"
	"github.com/aredoff/iptec/plugins/tor"
)

func main() {
	a := iptec.New()
	defer a.Close()
	// a.Use(asn.New())
	// a.Use(dns.New())
	a.Use(firehol.New())
	a.Use(bogon.New())
	a.Use(dnsbl.New())
	a.Use(tor.New())
	a.Activate()
	report, err := a.Find("1.1.1.1")
	// report, err := a.Find("2002:c000:200::1")
	// report, err := a.Find("192.42.116.195")
	// report, err := a.Find("2a0b:f4c2::29")
	// report, err = a.Find("2001:67c:6ec:203:192:42:116:180")
	// report, err = a.Find("2a0b:f4c2:2::38")
	// report, err = a.Find("2a0b:f4c2::25")
	// report, err = a.Find("2a0b:f4c1:2::251")
	if err != nil {
		fmt.Println("error")
	}
	reportString, err := json.Marshal(report)
	if err != nil {
		fmt.Println("eeer")
	}
	fmt.Println(string(reportString))
	// "something"
}
