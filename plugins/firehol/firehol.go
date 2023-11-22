package firehol

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"sync"
)

func newNetList() *netList {
	return &netList{
		Ips:  map[string]net.IP{},
		Nets: map[string]net.IPNet{},
	}
}

type netList struct {
	Ips  map[string]net.IP
	Nets map[string]net.IPNet
}

func (d *netList) AddIp(ip net.IP) {
	d.Ips[ip.String()] = ip
}

func (d *netList) AddNet(network net.IPNet) {
	d.Nets[network.String()] = network
}

func (d *netList) Find(ip net.IP) bool {
	_, ok := d.Ips[ip.String()]
	if ok {
		return true
	}
	for _, v := range d.Nets {
		if v.Contains(ip) {
			return true
		}
	}
	return false
}

func newFireholData() *fireholData {
	return &fireholData{
		lists: map[string]*netList{},
		mu:    &sync.Mutex{},
	}
}

type fireholData struct {
	lists map[string]*netList
	mu    *sync.Mutex
}

func (d *fireholData) List() []string {
	d.mu.Lock()
	defer d.mu.Unlock()
	res := []string{}
	for k, v := range d.lists {
		for _, ad := range v.Ips {
			res = append(res, fmt.Sprintf("%s:%s", k, ad.String()))
		}
		for _, ad := range v.Nets {
			res = append(res, fmt.Sprintf("%s:%s", k, ad.String()))
		}
	}
	return res
}

func (d *fireholData) AddIp(name string, ip net.IP) {
	d.mu.Lock()
	defer d.mu.Unlock()
	list, ok := d.lists[name]
	if !ok {
		d.lists[name] = newNetList()
		d.lists[name].AddIp(ip)
	} else {
		list.AddIp(ip)
	}
}

func (d *fireholData) AddNet(name string, network net.IPNet) {
	d.mu.Lock()
	defer d.mu.Unlock()
	list, ok := d.lists[name]
	if !ok {
		d.lists[name] = newNetList()
		d.lists[name].AddNet(network)
	} else {
		list.AddNet(network)
	}
}

func (d *fireholData) Find(ip net.IP) []string {
	d.mu.Lock()
	defer d.mu.Unlock()
	lists := make([]string, 0, len(d.lists))
	for name, list := range d.lists {
		if list.Find(ip) {
			lists = append(lists, name)
		}
	}
	return lists
}

func (d *fireholData) Serialization() ([]byte, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&d.lists)
	if err != nil {
		return nil, fmt.Errorf("cant serialize fireholData, err=%s", err)
	}
	return buffer.Bytes(), nil
}

func (d *fireholData) Deserialization(data []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&d.lists)
	if err != nil {
		return fmt.Errorf("cant deserialize fireholData, err=%s", err)
	}
	return nil
}
