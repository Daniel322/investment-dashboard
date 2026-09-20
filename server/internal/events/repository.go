package events

import (
	"context"
	"database/sql"
	"errors"
)

type EventsRepository struct {
	db *sql.DB
}

type EventFilter struct{}

const insertEvent = `
	INSERT INTO events (event, type, payload_id, created_at, data)
	VALUES ($1, $2, $3, $4, $5)
`

func (r *EventsRepository) Save(ctx context.Context, event *EventRecord) (*EventRecord, error) {
	if event == nil {
		return nil, ErrEventPointerIsNil
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, insertEvent, event.Event, event.Type, event.PayloadId, event.CreatedAt, event.Data)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventsRepository) Load(ctx context.Context, id string) (*EventRecord, error) {
	return nil, errors.New("not implemented")
}

func (r *EventsRepository) Find(ctx context.Context, filter EventFilter) (*EventRecord, error) {
	return nil, errors.New("not implemented")
}

func (r *EventsRepository) List(ctx context.Context, filter EventFilter) ([]*EventRecord, error) {
	return nil, errors.New("not implemented")
}
