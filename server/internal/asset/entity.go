package asset

import (
	"investment-dashboard/pkg/price"
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
	Price     *price.Price `json:"price"`
	Quantity  float64      `json:"quantity"`
	Type      string       `json:"type"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func Create(name string, slug string, type_ string) (*Asset, CreateAssetEvent) {
	asset := constructor(
		uuid.New(),
		name,
		slug,
		price.New(0),
		0,
		type_,
		time.Now(),
		time.Now(),
	)

	event := CreateAssetEvent{
		AssetID:   &asset.ID,
		CreatedAt: asset.CreatedAt,
	}

	return asset, event
}

func constructor(
	id uuid.UUID,
	name string,
	slug string,
	price *price.Price,
	quantity float64,
	type_ string,
	created_at time.Time,
	updated_at time.Time,
) *Asset {
	return &Asset{
		ID:        id,
		Name:      name,
		Slug:      slug,
		Price:     price,
		Quantity:  quantity,
		Type:      type_,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
}
