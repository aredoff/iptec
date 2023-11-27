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

func (p *Blacklist) loadLists(ss []*source) ([]*parsedSource, error) {
	var wg sync.WaitGroup
	ch := make(chan *parsedSource, len(sources))

	for _, v := range ss {
		data, err := p.Cash.Get(v.Name)
		if err == nil {
			lines, err := deserialization(data)
			if err != nil {
				return nil, err
			}
			ch <- &parsedSource{
				Lines:  lines,
				Source: v,
			}
			p.Log.Info(fmt.Sprintf("Load data from cash. [%s]", v.Name))
			continue
		}
		wg.Add(1)
		go func(s *source) {
			defer wg.Done()

			r, err := p.Client.Get(s.Url)
			if err != nil {
				p.Log.Error(fmt.Sprintf("Load data from cash. [%s] Error: %s", name, err))
				return
			}
			if r.StatusCode != 200 {
				p.Log.Error(fmt.Sprintf("Load data from cash. [%s] Status code=%d", name, r.StatusCode))
				return
			}
			parsed := &parsedSource{
				Lines:  s.extractor(string(r.Body)),
				Source: s,
			}
			data, err := serialization(parsed.Lines)
			if err != nil {
				p.Log.Error(fmt.Sprintf("cant serialization lines, err=%s", err))
				return
			}
			err = p.Cash.Set(s.Name, data)
			if err != nil {
				p.Log.Error(fmt.Sprintf("cant set in cash data, err=%s", err))
				return
			}

			p.Log.Info(fmt.Sprintf("Load data from source. [%s]", s.Name))
			ch <- parsed
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sourcesLists := []*parsedSource{}

	for parsedSources := range ch {
		exist, err := p.Cash.Exist(parsedSources.Source.Name)
		if err != nil {
			return nil, fmt.Errorf("cant check exist in cash data, err=%s", err)
		}
		if !exist {
			data, err := serialization(parsedSources.Lines)
			if err != nil {
				return nil, fmt.Errorf("cant serialization lines, err=%s", err)
			}
			err = p.Cash.Set(parsedSources.Source.Name, data)
			if err != nil {
				return nil, fmt.Errorf("cant set in cash data, err=%s", err)
			}
		}
		sourcesLists = append(sourcesLists, parsedSources)
	}
	return sourcesLists, nil
}

func (p *Blacklist) Activate() error {
	p.data = newBlacklist()
	parsedSources, err := p.loadLists(sources)
	if err != nil {
		return err
	}
	for _, s := range parsedSources {
		for _, line := range s.Lines {
			address := net.ParseIP(line)
			if address != nil {
				p.data.AddIp(address, s.Source)
			} else {
				_, network, err := net.ParseCIDR(line)
				if err != nil {
					continue
				}
				p.data.AddNet(*network, s.Source)
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

	a := blacklistResult{
		Lists:  srcslist,
		points: points,
	}
	return &a, nil
}
