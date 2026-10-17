package interfaces

type GetPriceResponse struct {
	Slug  string  `json:"slug"`
	Price float64 `json:"price"`
}
