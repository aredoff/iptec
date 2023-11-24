package tor

type torResult struct {
	Detect bool     `json:"detect"`
	Proofs []string `json:"proofs"`
}

func (r *torResult) Points() int {
	if r.Detect {
		return 10
	}
	return 0
}
