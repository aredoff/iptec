package iptec

import "time"

type appReport struct {
	Address string                  `json:"address"`
	Date    time.Time               `json:"time"`
	Plugins map[string]PluginReport `json:"plugins"`
	Points  int                     `json:"points"`
}
