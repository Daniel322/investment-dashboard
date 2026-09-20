package asset

import (
	"investment-dashboard/interfaces"
	"log"
)

// тут подписки на события связанные с обновлением assets

type eventBusSub interface {
	Subscribe(topic string, buffer int) (<-chan interfaces.Event[any], func())
}

type AssetEvents struct {
	EventBus eventBusSub
}

// добавить публичные обертки для событий (SubscribeAssetCreated)
// вызывать их в ините модуля
func (a *AssetEvents) subscribe(topic string, cb func(event interfaces.Event[any])) {
	ch, unsubscribe := a.EventBus.Subscribe(topic, 1000)

	defer unsubscribe()

	for ev := range ch {
		cb(ev)
	}
}

func asTypedHandler[T any](cb func(event interfaces.Event[T])) func(event interfaces.Event[any]) {
	return func(event interfaces.Event[any]) {
		payload, ok := event.Payload.(T)
		if !ok {
			log.Printf("handler: unexpected payload type for %s: %T", event.Topic, event.Payload)
			return
		}

		cb(interfaces.Event[T]{
			Topic:   event.Topic,
			Payload: payload,
		})
	}
}

type UpdatePriceEvent struct {
	AssetID string
	Price   float64
}

func (a *AssetEvents) SubscribeAssetPriceUpdated() {
	cb := func(event interfaces.Event[UpdatePriceEvent]) {
		log.Println("asset price updated:", event.Payload.AssetID, event.Payload.Price)
	}

	a.subscribe("asset.price-updated", asTypedHandler(cb))
}
