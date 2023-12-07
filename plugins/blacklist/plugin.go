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

func (p *Blacklist) loadList(s *source) ([]string, error) {
	data, err := p.Cash.Get(s.Name)
	if err == nil {
		lines, err := deserialization(data)
		if err != nil {
			return nil, err
		}
		// p.Log.Info(fmt.Sprintf("Load data from cash. [%s]", s.Name))
		return lines, nil
	}

	r, err := p.Client.Get(s.Url)
	if err != nil {
		return nil, fmt.Errorf("load data from cash. [%s] Error: %s", name, err)
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("load data from cash. [%s] Status code=%d", name, r.StatusCode)
	}

	lines := s.extractor(string(r.Body))

	data, err = serialization(lines)
	if err != nil {
		return nil, fmt.Errorf("cant serialization lines, err=%s", err)
	}
	err = p.Cash.Set(s.Name, data)
	if err != nil {
		return nil, fmt.Errorf("cant set in cash data, err=%s", err)
	}
	// p.Log.Info(fmt.Sprintf("Load data from source. [%s]", s.Name))
	return lines, nil
}

type parsedSource struct {
	Lines  []string
	Source *source
	Err    error
}

func (p *Blacklist) Activate() error {
	p.data = newBlacklist()

	var wg sync.WaitGroup
	ch := make(chan *parsedSource, len(sources))

	for _, v := range sources {
		wg.Add(1)
		go func(s *source) {
			defer wg.Done()

			lines, err := p.loadList(s)
			if err != nil {
				ch <- &parsedSource{
					Err:    err,
					Source: s,
				}
			}
			ch <- &parsedSource{
				Lines:  lines,
				Source: s,
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for parsedSources := range ch {
		if parsedSources.Err != nil {
			return parsedSources.Err
		}
		for _, line := range parsedSources.Lines {
			address := net.ParseIP(line)
			if address != nil {
				p.data.AddIp(address, parsedSources.Source)
			} else {
				_, network, err := net.ParseCIDR(line)
				if err != nil {
					continue
				}
				p.data.AddNet(*network, parsedSources.Source)
			}
		}
	}

	p.Log.Info(fmt.Sprintf("Data collected. [%d]", len(p.data.Ips)+len(p.data.Nets)))
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

	a := Report{
		Lists:  srcslist,
		points: points,
	}
	return &a, nil
}
