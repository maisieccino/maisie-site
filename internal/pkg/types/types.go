package types

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Path   string
	Method string
}

type ServerType interface {
	GetRouter() chi.Router
}

// RouteMap is a mapping of routes to a handler function.
// s is the server that provides data + logging for the handler.
type RouteMap[s ServerType] map[Route]func(s) http.HandlerFunc

// Build takes the chi router and sets up routes for each
// route defined in the map.
func (m RouteMap[T]) Build(s T) {
	for r, handler := range m {
		if r.Method == "*" {
			s.GetRouter().Mount(r.Path, handler(s))
		} else {
			s.GetRouter().Method(r.Method, r.Path, handler(s))
		}
	}
}
