package events

import (
	"time"

	"github.com/google/uuid"
)

type Event[T any] struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	AssetID   *uuid.UUID `json:"asset_id"`
	CreatedAt time.Time  `json:"created_at"`
	Data      T          `json:"data"`
}
