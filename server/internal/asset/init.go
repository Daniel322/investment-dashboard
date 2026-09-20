package asset

import (
	"database/sql"
)

type AssetModule struct {
	CreateAssetCommandHandler *CreateAssetCommandHandler
	FindAssetQueryHandler     *FindAssetQueryHandler
}

func Init(db *sql.DB) *AssetModule {
	repository := &AssetRepository{
		db: db,
	}
	CreateAssetCommandHandler := &CreateAssetCommandHandler{
		Repository: repository,
	}
	FindAssetQueryHandler := &FindAssetQueryHandler{
		Repository: repository,
	}

	return &AssetModule{
		CreateAssetCommandHandler: CreateAssetCommandHandler,
		FindAssetQueryHandler:     FindAssetQueryHandler,
	}
}
