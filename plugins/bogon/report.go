package bogon

type bogonResult struct {
	Bogon bool `json:"bogon"`
}

func (r *bogonResult) Points() int {
	if r.Bogon {
		return 100
	}
	return 0
}
