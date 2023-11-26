package blacklist

type blacklistResult struct {
	Lists  []string `json:"lists"`
	points int
}

func (r *blacklistResult) Points() int {
	return r.points
}
