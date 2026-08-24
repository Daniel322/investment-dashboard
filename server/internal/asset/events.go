package asset

import (
	"investment-dashboard/internal/events"
	"investment-dashboard/pkg/price"
	"time"

	"github.com/google/uuid"
)

type CreateAssetEventData struct {
	AssetID   uuid.UUID `json:"asset_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAssetEvent = events.Event[CreateAssetEventData]

type UpdateQuantityEventData struct {
	Quantity float64 `json:"quantity"`
}

type UpdateQuantityEvent = events.Event[UpdateQuantityEventData]

func (a *Asset) UpdateQuantity(event *UpdateQuantityEvent) {
	a.Quantity = event.Data.Quantity
}

type UpdatePriceEventData struct {
	Price *price.Price `json:"price"`
}

type UpdatePriceEvent = events.Event[UpdatePriceEventData]

func (a *Asset) UpdatePrice(event *UpdatePriceEvent) {
	a.Price = event.Data.Price
}
