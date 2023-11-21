package iptec

import (
	"github.com/aredoff/iptec/webclient"
)

type webclientMixinInterface interface {
	webclientInitialization()
}

type WebClient struct {
	Client *webclient.Client
}

func (m *WebClient) webclientInitialization() {
	m.Client = webclient.NewClient()
}
