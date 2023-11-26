package iptec

import "github.com/aredoff/iptec/dnsclient"

type dnsclientMixinInterface interface {
	dnsclientInitialization()
}

type DnsClient struct {
	Dns *dnsclient.Dns
}

func (m *DnsClient) dnsclientInitialization() {
	m.Dns = dnsclient.New()
}
