package asset

import (
	"context"
	"investment-dashboard/interfaces"
)

type FindAssetQueryHandler struct {
	Repository interfaces.Repository[AssetRecord, AssetFilter]
}

type FindAssetQuery struct {
	Name *string
	Slug *string
	Type *string
}

func (instance *FindAssetQueryHandler) Handle(
	ctx context.Context,
	query FindAssetQuery,
) (*AssetRecord, error) {
	filter := AssetFilter{
		Name: query.Name,
		Slug: query.Slug,
		Type: query.Type,
	}
	asset, err := instance.Repository.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return asset, err
}
