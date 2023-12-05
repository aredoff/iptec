package iptec

import (
	"fmt"
	"net"
	"sync"
	"time"

	clog "github.com/aredoff/iptec/log"
)

const Version = "0.0.5"

func New() *App {
	return &App{
		plugins:          map[string]Plugin{},
		activatedPlugins: map[string]Plugin{},
		cash:             NewCach(),
		activated:        false,
		mu:               &sync.Mutex{},
	}
}

type App struct {
	plugins          map[string]Plugin
	activatedPlugins map[string]Plugin
	cash             *cash
	activated        bool
	mu               *sync.Mutex
}

func (a *App) Use(p Plugin) {
	defer a.mu.Unlock()
	a.mu.Lock()
	a.plugins[p.Name()] = p

	curatorMixin, ok := p.(curatorMixinInterface)
	if ok {
		curatorMixin.curatorInitialization(p.Name(), a)
	}

	cachMixin, ok := p.(cashMixinInterface)
	if ok {
		cachMixin.cashInitialization(p.Name(), a.cash)
	}

	webclientMixin, ok := p.(webclientMixinInterface)
	if ok {
		webclientMixin.webclientInitialization()
	}

	dnsclientMixin, ok := p.(dnsclientMixinInterface)
	if ok {
		dnsclientMixin.dnsclientInitialization()
	}
}

func (a *App) Activate() {
	var wg sync.WaitGroup
	for _, v := range a.plugins {
		wg.Add(1)
		go func(plugin Plugin) {
			defer wg.Done()
			err := plugin.Activate()
			if err != nil {
				clog.Error(err)
				return
			}
			a.mu.Lock()
			a.activatedPlugins[plugin.Name()] = plugin
			a.mu.Unlock()
			clog.Info(fmt.Sprintf("plugin/%s activated", plugin.Name()))
		}(v)
	}
	wg.Wait()
	a.activated = true
}

func (a *App) Find(address string) (*appReport, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return nil, fmt.Errorf("cant read IP: %s", address)
	}
	if !a.activated {
		a.Activate()
	}
	defer a.mu.Unlock()
	a.mu.Lock()
	var wg sync.WaitGroup
	ch := make(chan pluginReport, len(a.activatedPlugins))
	for _, v := range a.activatedPlugins {
		wg.Add(1)
		go func(plugin Plugin) {
			defer wg.Done()
			rep, err := plugin.Find(ip)
			if err != nil {
				clog.Error(err)
				return
			}
			if rep.Points() != 0 {
				ch <- pluginReport{
					report: rep,
					name:   plugin.Name(),
				}
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	report := &appReport{
		Address: address,
		Date:    time.Now(),
		Plugins: map[string]PluginReport{},
	}
	points := 0
	for result := range ch {
		report.Plugins[result.name] = result.report
		points += result.report.Points()
	}
	report.Points = points
	return report, nil
}

func (a *App) Close() {
	a.cash.Close()
}
