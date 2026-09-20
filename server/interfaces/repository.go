package interfaces

import "context"

type Repository[Entity any, Filter any] interface {
	Save(ctx context.Context, entity *Entity) (*Entity, error)
	Load(ctx context.Context, id string) (*Entity, error)
	List(ctx context.Context, filter Filter) ([]*Entity, error)
	Find(ctx context.Context, filter Filter) (*Entity, error)
}
