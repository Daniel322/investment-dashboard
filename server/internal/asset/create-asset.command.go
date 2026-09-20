package asset

import (
	"context"
	"investment-dashboard/interfaces"
	"log"
)

type eventBus interface {
	Publish(ctx context.Context, event interfaces.Event[any])
}

type CreateAssetCommandHandler struct {
	Repository interfaces.Repository[AssetRecord, AssetFilter]
	EventBus   eventBus
}

type CreateAssetCommand struct {
	Name  string
	Slug  string
	Type  string
	Price *int
	Rate  *int
}

func (instance *CreateAssetCommandHandler) Handle(
	ctx context.Context,
	cmd CreateAssetCommand,
) (*AssetRecord, error) {
	assetInRepository, err := instance.Repository.Find(ctx, AssetFilter{
		Name: &cmd.Name,
		Slug: &cmd.Slug,
		Type: &cmd.Type,
	})
	if err != nil && err != ErrAssetNotFound {
		return nil, err
	}
	if assetInRepository != nil {
		log.Println("asset already exists in repository", assetInRepository.Name, assetInRepository.Slug, assetInRepository.Type)
		return assetInRepository, nil
	}

	asset, err := Create(cmd.Name, cmd.Slug, cmd.Type, cmd.Price, cmd.Rate)
	if err != nil {
		return nil, err
	}

	// save entity in repository
	saveResult, err := instance.Repository.Save(ctx, asset.ToRecord())

	if err != nil {
		return nil, err
	}

	// publish event
	instance.EventBus.Publish(ctx, interfaces.Event[any]{
		Topic:   "asset.created",
		Payload: saveResult,
	})

	return saveResult, nil
}
