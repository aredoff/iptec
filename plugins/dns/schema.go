package dns

type DnsResult struct {
	asn string `json:"asn"`
}

func (r *DnsResult) Points() int {
	return 1
}
