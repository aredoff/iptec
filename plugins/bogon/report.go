package bogon

type Report struct {
	Bogon bool `json:"bogon"`
}

func (r *Report) Points() int {
	if r.Bogon {
		return 100
	}
	return 0
}
