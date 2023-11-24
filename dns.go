package iptec

import "github.com/aredoff/iptec/dnslient"

type dnsclientMixinInterface interface {
	dnsclientInitialization()
}

type DnsClient struct {
	Dns *dnslient.Dns
}

func (m *DnsClient) dnsclientInitialization() {
	m.Dns = dnslient.New()
}
