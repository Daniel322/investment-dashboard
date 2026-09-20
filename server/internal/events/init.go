package events

import (
	"database/sql"
	"investment-dashboard/pkg/event_bus"
)

func Init(db *sql.DB, eventBus *event_bus.Bus) {
	repository := &EventsRepository{
		db: db,
	}
	events := &Events{
		Repository: repository,
		EventBus:   eventBus,
	}

	go events.SubscribeCreateAsset()
}
