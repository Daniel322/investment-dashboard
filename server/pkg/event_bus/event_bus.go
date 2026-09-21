package event_bus

import (
	"context"
	"investment-dashboard/interfaces"
	"log"
	"sync"
)

type Bus struct {
	mu   sync.RWMutex
	subs map[string]map[chan interfaces.Event[any]]struct{}
}

func NewBus() *Bus {
	return &Bus{subs: make(map[string]map[chan interfaces.Event[any]]struct{})}
}

func (b *Bus) SubscribeHandlers(handlers map[string]interfaces.EventHandler[any]) {
	for topic, handler := range handlers {
		go func(topic string, handler interfaces.EventHandler[any]) {
			ch, unsubscribe := b.subscribe(topic, 1000)
			defer unsubscribe()
			for ev := range ch {
				if err := handler.Handle(ev); err != nil {
					log.Println("error in event handler:", topic, err)
				}
			}
		}(topic, handler)
	}
}

func (b *Bus) subscribe(topic string, buffer int) (<-chan interfaces.Event[any], func()) {
	ch := make(chan interfaces.Event[any], buffer)

	b.mu.Lock()
	if b.subs[topic] == nil {
		b.subs[topic] = make(map[chan interfaces.Event[any]]struct{})
	}
	b.subs[topic][ch] = struct{}{}
	b.mu.Unlock()

	unsubscribe := func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		if subs, ok := b.subs[topic]; ok {
			if _, ok := subs[ch]; ok {
				delete(subs, ch)
				close(ch)
			}
			if len(subs) == 0 {
				delete(b.subs, topic)
			}
		}
	}

	return ch, unsubscribe
}

func (b *Bus) Publish(ctx context.Context, ev interfaces.Event[any]) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for ch := range b.subs[ev.Topic] {
		select {
		case ch <- ev:
		case <-ctx.Done():
			return
		default:
			log.Println("Buffer overflow", ev.Topic)
			// Буфер переполнен. В зависимости от задачи:
			// - дропнуть событие,
			// - залогировать,
			// - сделать блокирующую отправку,
			// - отправить в отдельную очередь retry.
		}
	}
}
