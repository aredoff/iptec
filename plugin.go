package iptec

import (
	"fmt"
)

type СuratorMixinInterface interface {
	curatorInitialization(app *App)
}

type СuratorMixin struct {
	app *App
}

func (m *СuratorMixin) curatorInitialization(app *App) {
	m.app = app
}

func (m *СuratorMixin) Plugin(pluginName string) (Plugin, error) {
	plugin, ok := m.app.plugins[pluginName]
	if ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("Plugin %s missing", pluginName)
}
