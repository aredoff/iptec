package dnsbl

type Report struct {
	Lists map[string]string `json:"lists"`
}

func (r *Report) Points() int {
	return len(r.Lists)
}
