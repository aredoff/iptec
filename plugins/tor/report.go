package tor

type Report struct {
	Detect bool     `json:"detect"`
	Proofs []string `json:"proofs"`
}

func (r *Report) Points() int {
	if r.Detect {
		return 10
	}
	return 0
}
