package dnsbl

type dnsblResult struct {
	Lists map[string]string `json:"lists"`
}

func (r *dnsblResult) Points() int {
	return len(r.Lists)
}
