package blacklist

import (
	"fmt"
	"net"
	"strings"
)

type ipSources struct {
	Ip      net.IP
	Sources []*source
}

func (ss *ipSources) AddSource(s *source) {
	for _, v := range ss.Sources {
		if s.Name == v.Name {
			return
		}
	}
	ss.Sources = append(ss.Sources, s)
}

type netSources struct {
	Net     net.IPNet
	Sources []*source
}

func (ss *netSources) AddSource(s *source) {
	for _, v := range ss.Sources {
		if s.Name == v.Name {
			return
		}
	}
	ss.Sources = append(ss.Sources, s)
}

func newBlacklist() *blacklist {
	return &blacklist{
		Ips:  map[string]*ipSources{},
		Nets: map[string]*netSources{},
	}
}

type blacklist struct {
	Ips  map[string]*ipSources
	Nets map[string]*netSources
}

func (l *blacklist) AddIp(ip net.IP, s *source) error {
	_, ok := l.Ips[ip.String()]
	if !ok {
		l.Ips[ip.String()] = &ipSources{
			Ip:      ip,
			Sources: []*source{},
		}
		l.Ips[ip.String()].AddSource(s)
	} else {
		l.Ips[ip.String()].AddSource(s)
	}
	return nil
}

func (l *blacklist) AddNet(network net.IPNet, s *source) error {
	_, ok := l.Nets[network.String()]
	if !ok {
		l.Nets[network.String()] = &netSources{
			Net:     network,
			Sources: []*source{},
		}
		l.Nets[network.String()].AddSource(s)
	} else {
		l.Nets[network.String()].AddSource(s)
	}
	return nil
}

func (l *blacklist) Find(ip net.IP) []*source {
	ipSrc, ok := l.Ips[ip.String()]
	if ok {
		return ipSrc.Sources
	}
	for _, v := range l.Nets {
		if v.Net.Contains(ip) {
			return v.Sources
		}
	}
	return []*source{}
}

func (l *blacklist) List() []string {
	res := []string{}

	for _, v := range l.Ips {
		sourcesNames := []string{}
		for _, s := range v.Sources {
			sourcesNames = append(sourcesNames, s.Name)
		}
		res = append(res, fmt.Sprintf("%s:%s", strings.Join(sourcesNames, ", "), v.Ip.String()))
	}
	for _, v := range l.Nets {
		sourcesNames := []string{}
		for _, s := range v.Sources {
			sourcesNames = append(sourcesNames, s.Name)
		}
		res = append(res, fmt.Sprintf("%s:%s", strings.Join(sourcesNames, ", "), v.Net.String()))
	}

	return res
}
