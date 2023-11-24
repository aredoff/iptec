package tor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"sync"
)

func newTorList() *torList {
	return &torList{
		Ips: map[string]net.IP{},
	}
}

type torList struct {
	Ips map[string]net.IP
}

func (d *torList) AddIp(ip net.IP) {
	d.Ips[ip.String()] = ip
}

func (d *torList) Find(ip net.IP) bool {
	_, ok := d.Ips[ip.String()]
	if ok {
		return true
	}
	return false
}

func newTorData() *torData {
	return &torData{
		lists: map[string]*torList{},
		mu:    &sync.Mutex{},
	}
}

type torData struct {
	lists map[string]*torList
	mu    *sync.Mutex
}

func (d *torData) List() []string {
	d.mu.Lock()
	defer d.mu.Unlock()
	res := []string{}
	for k, v := range d.lists {
		for _, ad := range v.Ips {
			res = append(res, fmt.Sprintf("%s:%s", k, ad.String()))
		}
	}
	return res
}

func (d *torData) AddIp(name string, ip net.IP) {
	d.mu.Lock()
	defer d.mu.Unlock()
	list, ok := d.lists[name]
	if !ok {
		d.lists[name] = newTorList()
		d.lists[name].AddIp(ip)
	} else {
		list.AddIp(ip)
	}
}

func (d *torData) Find(ip net.IP) []string {
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

func (d *torData) Serialization() ([]byte, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&d.lists)
	if err != nil {
		return nil, fmt.Errorf("cant serialize tor data, err=%s", err)
	}
	return buffer.Bytes(), nil
}

func (d *torData) Deserialization(data []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&d.lists)
	if err != nil {
		return fmt.Errorf("cant deserialize tor data, err=%s", err)
	}
	return nil
}
