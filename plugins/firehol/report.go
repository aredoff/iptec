package firehol

type fireholResult struct {
	Name  string   `json:"name"`
	Lists []string `json:"lists"`
}

func (r *fireholResult) Points() int {
	return len(r.Lists) * 2
}
