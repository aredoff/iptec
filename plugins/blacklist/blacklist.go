package blacklist

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

type ipSources struct {
	Ip      net.IP
	Sources []*source
}

func (ss ipSources) AddSource(s *source) {
	var find bool
	for _, v := range ss.Sources {
		if s.Name == v.Name {
			find = true
			break
		}
	}
	// ss.Sources = append(ss.Sources, s)
	if !find {
		ss.Sources = append(ss.Sources, s)
	}
}

type netSources struct {
	Net     net.IPNet
	Sources []*source
}

func (ss netSources) AddSource(s *source) {
	var find bool
	for _, v := range ss.Sources {
		if s.Name == v.Name {
			find = true
			break
		}
	}
	if !find {
		ss.Sources = append(ss.Sources, s)
	}
}

func newBlacklist(sources []source) *blacklist {
	newSources := map[string]source{}
	for _, v := range sources {
		newSources[v.Name] = v
	}
	return &blacklist{
		Ips:     map[string]ipSources{},
		Nets:    map[string]netSources{},
		Sources: newSources,
	}
}

type blacklist struct {
	Ips     map[string]ipSources
	Nets    map[string]netSources
	Sources map[string]*source
}

func (l *blacklist) AddIp(ip net.IP, sourceName string) error {

	_, ok := l.Sources[sourceName]
	if !ok {
		return fmt.Errorf("source %s dont exist", sourceName)
	}

	_, ok = l.Ips[ip.String()]
	if !ok {
		l.Ips[ip.String()] = ipSources{
			Ip:      ip,
			Sources: []*source{},
		}
		l.Ips[ip.String()].AddSource(&l.Sources[sourceName])
	} else {
		l.Ips[ip.String()].AddSource(&s)
	}
	return nil
}

func (l *blacklist) AddNet(network net.IPNet, sourceName string) error {
	s, ok := l.Sources[sourceName]
	if !ok {
		return fmt.Errorf("source %s dont exist", sourceName)
	}

	_, ok = l.Nets[network.String()]
	if !ok {
		l.Nets[network.String()] = netSources{
			Net:     network,
			Sources: []*source{&s},
		}
	} else {
		l.Nets[network.String()].AddSource(&s)
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

func (l *blacklist) Serialization() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&l)
	if err != nil {
		return nil, fmt.Errorf("cant serialize blacklist, err=%s", err)
	}
	return buffer.Bytes(), nil
}

func (l *blacklist) Deserialization(data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&l)
	if err != nil {
		return fmt.Errorf("cant deserialize blacklist, err=%s", err)
	}
	return nil
}
