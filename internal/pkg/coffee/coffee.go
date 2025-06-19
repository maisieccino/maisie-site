// Package coffee is a server for the London Coffee Map.
package coffee

import "context"

type Store interface {
	Get(context.Context, string) (MapItem, error)
	Create(context.Context, MapItem) error
	Update(context.Context, string, MapItem) error
	List(context.Context) ([]MapItem, error)
	Delete(context.Context, string) error
}
