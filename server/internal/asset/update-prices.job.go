package asset

import (
	"context"
	"investment-dashboard/interfaces"
	"investment-dashboard/pkg/asset_type"
	"log"
)

type Adapter interface {
	GetPrice(slug string) (*interfaces.GetPriceResponse, error)
}

type UpdatePricesJob struct {
	adapter  Adapter
	type_    asset_type.AssetType
	handler  GetAssetsByTypeQueryHandler
	eventBus eventBus
}

type UpdatePricesEvent struct {
	AssetID string
	Price   float64
}

func NewUpdatePricesJob(
	adapter Adapter,
	type_ asset_type.AssetType,
	handler GetAssetsByTypeQueryHandler,
	eventBus eventBus,
) *UpdatePricesJob {
	return &UpdatePricesJob{
		adapter:  adapter,
		type_:    type_,
		handler:  handler,
		eventBus: eventBus,
	}
}

func (job *UpdatePricesJob) Run() error {
	log.Println("UpdatePricesJob.Run")
	ctx := context.Background()
	jobTypeStr := job.type_.String()
	assets, err := job.handler.Handle(ctx, GetAssetsByTypeQuery{
		Type: &jobTypeStr,
	})
	if err != nil {
		return err
	}
	for _, asset := range assets {
		price, err := job.adapter.GetPrice(asset.Slug)
		if err != nil {
			return err
		}
		job.eventBus.Publish(ctx, interfaces.Event[any]{
			Topic: "asset.price-updated",
			Payload: UpdatePricesEvent{
				AssetID: asset.ID,
				Price:   price.Price,
			},
		})
	}
	return nil
}
