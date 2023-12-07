package blacklist

type Report struct {
	Lists  []string `json:"lists"`
	points int
}

func (r *Report) Points() int {
	return r.points
}
