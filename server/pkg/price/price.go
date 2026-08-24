package price

type Price struct {
	Minor int `json:"minor"`
}

func (p *Price) Price() float64 {
	return float64(p.Minor) / 100
}

func New(minor int) *Price {
	return &Price{Minor: minor}
}
