package dnslient

import (
	"context"
	"net"
	"time"
)

func New() *Dns {
	return &Dns{
		timeout: 4 * time.Second,
	}
}

type Dns struct {
	client  net.Resolver
	timeout time.Duration
}

func (c *Dns) A(host string) (records []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	records, err = c.client.LookupHost(ctx, host)
	return
}

func (c *Dns) TXT(host string) (records []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	records, err = c.client.LookupTXT(ctx, host)
	return
}

func (c *Dns) CNAME(host string) (records string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	records, err = c.client.LookupCNAME(ctx, host)
	return
}

func (c *Dns) MX(host string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	mxs, err := c.client.LookupMX(ctx, host)
	if err != nil {
		return nil, err
	}
	records := make([]string, 0, len(mxs))
	for _, v := range mxs {
		records = append(records, v.Host)
	}
	return records, nil
}
