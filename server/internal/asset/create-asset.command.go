package asset

import (
	"context"
	"investment-dashboard/interfaces"
)

// type EventBus interface {
// 	Send(ctx context.Context, name string, event any) error
// }

type CreateAssetCommandHandler struct {
	Repository interfaces.Repository[AssetRecord, AssetFilter]
	// Eventbus   EventBus
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
	// create entity and validate
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
	// err = instance.Eventbus.Send(ctx, "asset.created", saveResult)

	return saveResult, nil
}
