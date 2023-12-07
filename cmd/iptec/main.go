package main

import (
	"encoding/json"
	"fmt"

	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/blacklist"
	"github.com/aredoff/iptec/plugins/tor"
)

func main() {
	a := iptec.New()
	defer a.Close()
	// a.Use(asn.New())
	// a.Use(dns.New())
	// a.Use(firehol.New())
	// a.Use(bogon.New())
	// a.Use(dnsbl.New())
	a.Use(tor.New())
	a.Use(blacklist.New())
	a.Activate()
	report, err := a.Find("204.137.14.106")
	// report, err := a.Find("2002:c000:200::1")
	// report, err := a.Find("192.42.116.195")
	// report, err := a.Find("2a0b:f4c2::29")
	// report, err = a.Find("2001:67c:6ec:203:192:42:116:180")
	// report, err = a.Find("2a0b:f4c2:2::38")
	// report, err = a.Find("2a0b:f4c2::25")
	// report, err = a.Find("2a0b:f4c1:2::251")

	for _, v := range report.Plugins {
		switch r := v.(type) {
		case *blacklist.Report:
			fmt.Println(r.Lists)
		case *tor.Report:
			fmt.Println(r.Detect)
		}
		// fmt.Println(v)
	}

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
