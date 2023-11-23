package main

import (
	"encoding/json"
	"fmt"

	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/dnsbl"
)

func main() {
	a := iptec.New()
	defer a.Close()
	// a.Use(asn.New())
	// a.Use(dns.New())
	// a.Use(firehol.New())
	// a.Use(bogon.New())
	a.Use(dnsbl.New("8.8.8.8"))
	a.Activate()
	report, err := a.Find("2002:c000:200::1")
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
