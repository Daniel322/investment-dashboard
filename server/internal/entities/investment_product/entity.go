package investment_product

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type InvestmentProduct struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Price_minor int       `json:"price_minor"`
	Quantity    float64   `json:"quantity"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Create(
	name string,
	slug string,
	type_ string,
) *InvestmentProduct {
	return constructor(
		uuid.New().String(),
		name,
		slug,
		0,
		0.00,
		type_,
		time.Now(),
		time.Now(),
	)
}

func constructor(
	id string,
	name string,
	slug string,
	price_minor int,
	quantity float64,
	type_ string,
	created_at time.Time,
	updated_at time.Time,
) *InvestmentProduct {
	return &InvestmentProduct{
		ID:          id,
		Name:        name,
		Slug:        slug,
		Price_minor: price_minor,
		Quantity:    quantity,
		Type:        type_,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
}

func (i *InvestmentProduct) GetPrice() float64 {
	return float64(i.Price_minor) / 100
}

func (i *InvestmentProduct) GetPriceString() string {
	return fmt.Sprintf("%d", i.Price_minor)
}

func (i *InvestmentProduct) GetQuantity() float64 {
	return float64(i.Quantity)
}

func (i *InvestmentProduct) GetQuantityString() string {
	return fmt.Sprintf("%f", i.Quantity)
}

func (i *InvestmentProduct) SetPrice(price float64) {
	i.Price_minor = int(price * 100)
}

func (i *InvestmentProduct) SetQuantity(quantity float64) {
	i.Quantity = quantity
}
