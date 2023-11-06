package main

import (
	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/asn"
)

func main() {
	a := iptec.New()
	defer a.Close()
	a.Use(asn.New())
	a.Activate()
	a.Collect()
}
