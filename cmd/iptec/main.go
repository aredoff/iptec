package main

import (
	"encoding/json"
	"fmt"

	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/firehol"
)

func main() {
	a := iptec.New()
	defer a.Close()
	// a.Use(asn.New())
	// a.Use(dns.New())
	a.Use(firehol.New())
	a.Activate()
	report, err := a.Find("36.0.8.1")
	if err != nil {
		fmt.Println("error")
	}
	reportString, err := json.Marshal(report)
	if err != nil {
		fmt.Println("eeer")
	}
	fmt.Println(string(reportString))
	// "something"
	fmt.Println(report)
}
