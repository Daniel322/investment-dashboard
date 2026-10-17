package events

import (
	"database/sql"
)

type EventsModuleCommands struct{}
type EventsModuleQueries struct{}
type EventsModuleHandlers struct {
	CreateAssetEventHandler  *CreateAssetEventHandler
	UpdatePricesEventHandler *UpdatePricesEventHandler
}

type EventsModule struct {
	Commands EventsModuleCommands
	Queries  EventsModuleQueries
	Handlers EventsModuleHandlers
}

func Init(db *sql.DB) *EventsModule {
	repository := &EventsRepository{
		db: db,
	}
	createAssetEventHandler := &CreateAssetEventHandler{
		Repository: repository,
	}
	updatePricesEventHandler := &UpdatePricesEventHandler{
		Repository: repository,
	}
	return &EventsModule{
		Commands: EventsModuleCommands{},
		Queries:  EventsModuleQueries{},
		Handlers: EventsModuleHandlers{
			CreateAssetEventHandler:  createAssetEventHandler,
			UpdatePricesEventHandler: updatePricesEventHandler,
		},
	}
}
