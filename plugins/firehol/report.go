package firehol

type fireholResult struct {
	Lists []string `json:"lists"`
}

func (r *fireholResult) Points() int {
	return len(r.Lists) * 2
}
