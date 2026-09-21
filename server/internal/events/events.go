package events

import (
	"context"
	"investment-dashboard/interfaces"
	"investment-dashboard/internal/asset"
	"log"

	"github.com/google/uuid"
)

type CreateAssetEventHandler struct {
	Repository interfaces.Repository[EventRecord, EventFilter]
}

type CreateAssetEvent = *asset.AssetRecord

func (a *CreateAssetEventHandler) Handle(_event interfaces.Event[CreateAssetEvent]) error {
	log.Println("createAssetCb", _event)
	payId, err := uuid.Parse(_event.Payload.ID)
	log.Println("err Parse", err)
	if err != nil {
		return err
	}
	event, err := Create("asset.created", "assets", &payId, _event.Payload)
	log.Println("err Create", err)
	if err != nil {
		return err
	}

	_, err = a.Repository.Save(context.Background(), event.ToRecord())
	log.Println("err Save", err)
	if err != nil {
		return err
	}

	return nil
}
