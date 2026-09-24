package asset

import (
	"context"
	"investment-dashboard/interfaces"
)

type GetAssetsByTypeQueryHandler struct {
	Repository interfaces.Repository[AssetRecord, AssetFilter]
}

type GetAssetsByTypeQuery struct {
	Type *string
}

func (instance *GetAssetsByTypeQueryHandler) Handle(
	ctx context.Context,
	query GetAssetsByTypeQuery,
) ([]*AssetRecord, error) {
	filter := AssetFilter{
		Type: query.Type,
	}
	assets, err := instance.Repository.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return assets, err
}
