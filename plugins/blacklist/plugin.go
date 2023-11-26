package blacklist

import (
	"fmt"
	"net"
	"sync"

	"github.com/aredoff/iptec"
)

const (
	name = "blacklist"
)

func New() *Blacklist {
	return &Blacklist{
		name: name,
	}
}

type Blacklist struct {
	name string
	iptec.Сurator
	iptec.Cash
	iptec.WebClient
	data *blacklist
}

func (p *Blacklist) Name() string {
	return p.name
}

type parsedSource struct {
	Lines  []string
	Source *source
}

func (p *Blacklist) Activate() error {
	p.data = newBlacklist(sources)
	data, err := p.Cash.Get("blacklist")
	if err == nil {
		err := p.data.Deserialization(data)
		if err != nil {
			return err
		}
		p.Log.Info(fmt.Sprintf("Load data from cash. [%d]", len(p.data.Ips)+len(p.data.Nets)))
		return nil
	}
	var wg sync.WaitGroup
	ch := make(chan *parsedSource, len(sources))

	for _, v := range sources {
		wg.Add(1)
		go func(s source) {
			defer wg.Done()

			r, err := p.Client.Get(s.Url)
			if err != nil {
				p.Log.Error(fmt.Sprintf("cant load %s, err=%s", name, err))
				return
			}
			if r.StatusCode != 200 {
				p.Log.Error(fmt.Sprintf("cant load %s, status_code=%d", name, r.StatusCode))
				return
			}
			ch <- &parsedSource{
				Lines:  s.extractor(string(r.Body)),
				Source: &s,
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for parsedSources := range ch {
		for _, line := range parsedSources.Lines {
			address := net.ParseIP(line)
			if address != nil {
				p.data.AddIp(address, parsedSources.Source.Name)
			} else {
				_, network, err := net.ParseCIDR(line)
				if err != nil {
					continue
				}
				p.data.AddNet(*network, parsedSources.Source.Name)
			}
		}
	}
	for _, v := range p.data.List() {
		fmt.Println(v)
	}
	p.Log.Info(fmt.Sprintf("Collected data from sources. [%d]", len(p.data.Ips)+len(p.data.Nets)))

	data, err = p.data.Serialization()
	if err != nil {
		return err
	}
	err = p.Cash.Set("blacklist1", data)
	if err != nil {
		return fmt.Errorf("cant set in cash data, err=%s", err)
	}
	return nil
}

func (p *Blacklist) Find(address net.IP) (iptec.PluginReport, error) {
	list := p.data.Find(address)

	points := 0
	srcslist := []string{}
	for _, v := range list {
		points += v.Points
		srcslist = append(srcslist, v.Name)
	}

	a := blacklistResult{
		Lists:  srcslist,
		points: points,
	}
	return &a, nil
}
