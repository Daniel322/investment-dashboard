package asset

import (
	"database/sql"
	"investment-dashboard/pkg/event_bus"
)

type AssetModule struct {
	CreateAssetCommandHandler *CreateAssetCommandHandler
	FindAssetQueryHandler     *FindAssetQueryHandler
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
	// AssetEvents := &AssetEvents{
	// 	EventBus: eventBus,
	// }

	return &AssetModule{
		CreateAssetCommandHandler: CreateAssetCommandHandler,
		FindAssetQueryHandler:     FindAssetQueryHandler,
	}
}
