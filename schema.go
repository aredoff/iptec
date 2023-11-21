package iptec

import "time"

type appReport struct {
	address string        `json:"address"`
	date    time.Time     `json:"time"`
	plugins []interface{} `json:"plugins"`
}
