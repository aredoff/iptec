package iptec

import (
	"sync"
	"time"

	clog "github.com/aredoff/iptec/log"
)

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
		}(v)
	}
	wg.Wait()
	a.activated = true
}

func (a *App) Find(address string) (*appReport, error) {
	if !a.activated {
		a.Activate()
	}
	defer a.mu.Unlock()
	a.mu.Lock()
	var wg sync.WaitGroup
	ch := make(chan interface{}, len(a.activatedPlugins))
	for _, v := range a.activatedPlugins {
		wg.Add(1)
		go func(plugin Plugin) {
			defer wg.Done()
			res, err := plugin.Find(address)
			if err != nil {
				clog.Error(err)
				return
			}
			ch <- res
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	report := &appReport{
		address: address,
		date:    time.Now(),
		plugins: []interface{}{},
	}

	for result := range ch {
		report.plugins = append(report.plugins, result)
	}
	return report, nil
}

func (a *App) Close() {
	a.cash.Close()
}
