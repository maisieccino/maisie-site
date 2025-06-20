// Package types defines shared types amongst the servers.
package types

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const MethodWildcard = "*"

// Route is a tuple of URL path and HTTP method to route server requests to.
type Route struct {
	Path   string
	Method string
}

// ServerType defines common methods that a server should
// expose to run properly.
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
		if r.Method == MethodWildcard {
			s.GetRouter().Mount(r.Path, handler(s))
		} else {
			s.GetRouter().Method(r.Method, r.Path, handler(s))
		}
	}
}
