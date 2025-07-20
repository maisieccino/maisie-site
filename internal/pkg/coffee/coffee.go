// Package coffee is a server for the London Coffee Map.
// It offers a server to be able to get, read, create, and update map data.
package coffee

import "context"

// Store is an interface to storage for the coffee map data.
type Store interface {
	Get(context.Context, string) (MapItem, error)
	Create(context.Context, MapItem) error
	Update(context.Context, string, MapItem) error
	List(context.Context) ([]MapItem, error)
	Delete(context.Context, string) error
}
