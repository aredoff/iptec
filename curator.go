package iptec

import (
	"fmt"

	clog "github.com/aredoff/iptec/log"
)

type curatorMixinInterface interface {
	curatorInitialization(string, *App)
}

type Сurator struct {
	app *App
	Log clog.P
}

func (m *Сurator) curatorInitialization(pluginName string, app *App) {
	m.app = app
	m.Log = clog.NewWithPlugin(pluginName)
}

func (m *Сurator) Plugin(pluginName string) (Plugin, error) {
	plugin, ok := m.app.plugins[pluginName]
	if ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("Plugin %s missing", pluginName)
}
