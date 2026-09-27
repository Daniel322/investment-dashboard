package asset

import (
	"context"
	"investment-dashboard/interfaces"
	"investment-dashboard/pkg/price"
)

type UpdatePricesEventHandler struct {
	Repository interfaces.Repository[AssetRecord, AssetFilter]
}

func (a *UpdatePricesEventHandler) Handle(_event interfaces.Event[UpdatePricesEvent]) error {
	ctx := context.Background()
	assetRecord, err := a.Repository.Find(ctx, AssetFilter{
		ID: &_event.Payload.AssetID,
	})
	if err != nil {
		return err
	}
	asset, err := FromRecord(assetRecord)
	if err != nil {
		return err
	}
	priceMinor := price.ToMinor(_event.Payload.Price)
	asset.Price = price.New(priceMinor)
	assetRecord, err = a.Repository.Save(ctx, asset.ToRecord())
	if err != nil {
		return err
	}
	return nil
}
