package coffee

import "context"

type Store interface {
	Get(context.Context, int) (MapItem, error)
	Create(context.Context, MapItem) error
	Update(context.Context, int, MapItem) error
	List(context.Context) ([]MapItem, error)
	Delete(context.Context, int) error
}
