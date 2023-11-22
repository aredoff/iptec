package asn

type AsnResult struct {
	asn string `json:"asn"`
}

func (r *AsnResult) Points() int {
	return 1
}
