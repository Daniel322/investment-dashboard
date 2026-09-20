package rate

type Rate struct {
	Minor int `json:"minor"`
}

func (r *Rate) Rate() float64 {
	return float64(r.Minor) / 100
}

func ToMinor(rate float64) int {
	return int(rate * 100)
}

func New(minor int) *Rate {
	return &Rate{Minor: minor}
}
