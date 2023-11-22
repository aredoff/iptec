package iptec

import "net"

type Plugin interface {
	Name() string
	Activate() error
	Find(net.IP) (PluginReport, error)
}

type PluginReport interface {
	Points() int
}

type pluginReport struct {
	report PluginReport
	name   string
}
