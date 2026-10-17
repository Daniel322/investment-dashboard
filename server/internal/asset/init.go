package asset

import (
	"database/sql"
	"investment-dashboard/pkg/asset_type"
	"investment-dashboard/pkg/event_bus"
)

type AssetModuleCommands struct {
	CreateAssetCommandHandler *CreateAssetCommandHandler
}
type AssetModuleQueries struct {
	FindAssetQueryHandler       *FindAssetQueryHandler
	GetAssetsByTypeQueryHandler *GetAssetsByTypeQueryHandler
}
type AssetModuleHandlers struct {
	UpdatePricesEventHandler *UpdatePricesEventHandler
}

type AssetModuleJobs struct {
	UpdateCryptoPricesJob *UpdatePricesJob
}
type AssetModule struct {
	Commands AssetModuleCommands
	Queries  AssetModuleQueries
	Handlers AssetModuleHandlers
	Jobs     AssetModuleJobs
}

func Init(db *sql.DB, eventBus *event_bus.Bus, adapter Adapter) *AssetModule {
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
	UpdatePricesEventHandler := &UpdatePricesEventHandler{
		Repository: repository,
	}
	UpdateCryptoPricesJob := NewUpdatePricesJob(
		adapter,
		asset_type.Crypto,
		*GetAssetsByTypeQueryHandler,
		eventBus,
	)
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
		Handlers: AssetModuleHandlers{
			UpdatePricesEventHandler: UpdatePricesEventHandler,
		},
		Jobs: AssetModuleJobs{
			UpdateCryptoPricesJob: UpdateCryptoPricesJob,
		},
	}
}
