package asset

import (
	"database/sql"
	"investment-dashboard/pkg/event_bus"
)

type AssetModuleCommands struct {
	CreateAssetCommandHandler *CreateAssetCommandHandler
}
type AssetModuleQueries struct {
	FindAssetQueryHandler       *FindAssetQueryHandler
	GetAssetsByTypeQueryHandler *GetAssetsByTypeQueryHandler
}
type AssetModuleHandlers struct{}
type AssetModule struct {
	Commands AssetModuleCommands
	Queries  AssetModuleQueries
	Handlers AssetModuleHandlers
}

func Init(db *sql.DB, eventBus *event_bus.Bus) *AssetModule {
	repository := &AssetRepository{
		db: db,
	}
	CreateAssetCommandHandler := &CreateAssetCommandHandler{
		Repository: repository,
		EventBus:   eventBus,
	}
	FindAssetQueryHandler := &FindAssetQueryHandler{
		Repository: repository,
	}
	GetAssetsByTypeQueryHandler := &GetAssetsByTypeQueryHandler{
		Repository: repository,
	}
	// AssetEvents := &AssetEvents{
	// 	EventBus: eventBus,
	// }

	return &AssetModule{
		Commands: AssetModuleCommands{
			CreateAssetCommandHandler: CreateAssetCommandHandler,
		},
		Queries: AssetModuleQueries{
			FindAssetQueryHandler:       FindAssetQueryHandler,
			GetAssetsByTypeQueryHandler: GetAssetsByTypeQueryHandler,
		},
		Handlers: AssetModuleHandlers{},
	}
}
