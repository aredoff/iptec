package iptec

type Plugin interface {
	Name() string
	Activate() error
	Find(string) (interface{}, error)
}
