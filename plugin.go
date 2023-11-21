package iptec

import (
	"fmt"

	clog "github.com/aredoff/iptec/log"
)

type Plugin interface {
	Name() string
	Activate() error
	Find(string) (interface{}, error)
}

type curatorMixinInterface interface {
	curatorInitialization(string, *App)
}

type СuratorMixin struct {
	app *App
	Log clog.P
}

func (m *СuratorMixin) curatorInitialization(pluginName string, app *App) {
	m.app = app
	m.Log = clog.NewWithPlugin(pluginName)
}

func (m *СuratorMixin) Plugin(pluginName string) (Plugin, error) {
	plugin, ok := m.app.plugins[pluginName]
	if ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("Plugin %s missing", pluginName)
}
