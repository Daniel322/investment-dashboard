package interfaces

import "context"

type Result = map[string]Position

type Adapter interface {
	Balance(ctx context.Context) (Result, error)
}
