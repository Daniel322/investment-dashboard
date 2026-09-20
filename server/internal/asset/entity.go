package asset

import (
	"investment-dashboard/pkg/asset_type"
	"investment-dashboard/pkg/price"
	"investment-dashboard/pkg/rate"
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID        uuid.UUID            `json:"id"`
	Name      string               `json:"name"`
	Slug      string               `json:"slug"`
	Price     *price.Price         `json:"price"`
	Rate      *rate.Rate           `json:"rate"`
	Type      asset_type.AssetType `json:"type"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type AssetRecord struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Price     int       `json:"price"`
	Rate      int       `json:"rate"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(name string, slug string, type_ string, price_ *int, rate_ *int) (*Asset, error) {
	if price_ == nil {
		price_ = new(int)
	}
	if rate_ == nil {
		rate_ = new(int)
	}
	asset, err := constructor(
		uuid.New(),
		name,
		slug,
		price.New(*price_),
		rate.New(*rate_),
		type_,
		time.Now(),
		time.Now(),
	)

	return asset, err
}

func (a *Asset) ToRecord() *AssetRecord {

	return &AssetRecord{
		ID:        a.ID.String(),
		Name:      a.Name,
		Slug:      a.Slug,
		Price:     a.Price.Minor,
		Rate:      a.Rate.Minor,
		Type:      a.Type.String(),
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
func constructor(
	id uuid.UUID,
	name string,
	slug string,
	price *price.Price,
	rate *rate.Rate,
	type_ string,
	created_at time.Time,
	updated_at time.Time,
) (*Asset, error) {
	assetType, err := asset_type.New(type_)
	if err != nil {
		return nil, err
	}
	return &Asset{
		ID:        id,
		Name:      name,
		Rate:      rate,
		Slug:      slug,
		Price:     price,
		Type:      assetType,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}
