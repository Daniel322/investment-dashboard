package events

import (
	"context"
	"errors"
	"investment-dashboard/interfaces"
	"investment-dashboard/internal/asset"
	"log"

	"github.com/google/uuid"
)

type eventBusSub interface {
	Subscribe(topic string, buffer int) (<-chan interfaces.Event[any], func())
}

type Events struct {
	EventBus   eventBusSub
	Repository interfaces.Repository[EventRecord, EventFilter]
}

func (a *Events) subscribe(topic string, cb func(event interfaces.Event[any]) error) {
	ch, unsubscribe := a.EventBus.Subscribe(topic, 1000)

	defer unsubscribe()

	for ev := range ch {
		err := cb(ev)
		if err != nil {
			log.Println("error in event handler:", topic, err)
		}
	}
}

type CreateAssetEvent = *asset.AssetRecord

func (a *Events) createAssetCb(_event interfaces.Event[CreateAssetEvent]) error {
	log.Println("createAssetCb", _event)
	payId, err := uuid.Parse(_event.Payload.ID)
	log.Println("err Parse", err)
	if err != nil {
		return err
	}
	event, err := Create[CreateAssetEvent]("asset.created", "assets", &payId, _event.Payload)
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

func asTypedHandler[T any](cb func(event interfaces.Event[T]) error) func(event interfaces.Event[any]) error {
	return func(event interfaces.Event[any]) error {
		payload, ok := event.Payload.(T)
		if !ok {
			log.Printf("handler: unexpected payload type for %s: %T", event.Topic, event.Payload)
			return errors.New("unexpected payload type")
		}

		return cb(interfaces.Event[T]{
			Topic:   event.Topic,
			Payload: payload,
		})
	}
}

func (a *Events) SubscribeCreateAsset() {
	a.subscribe("asset.created", asTypedHandler(a.createAssetCb))
}
