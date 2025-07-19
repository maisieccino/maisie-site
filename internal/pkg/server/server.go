// Package server provides the root server for maisie's website.
package server

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/maisieccino/maisie-site/internal/pkg/coffee"
	"github.com/maisieccino/maisie-site/internal/pkg/middleware"
	"github.com/maisieccino/maisie-site/internal/pkg/types"
	"go.uber.org/zap"
)

type (
	DB struct {
		Enabled  bool   `keyval:"enabled"`
		Hostname string `keyval:"hostname"`
		Port     int    `keyval:"port"`
		User     string `keyval:"user"`
		Password string `keyval:"password"`
		Database string `keyval:"database"`
		Conn     *pgx.Conn
	}
	Config struct {
		Host       string `keyval:"host"`
		Port       int    `keyval:"port"`
		StaticPath string `keyval:"staticPath"`
		Logger     *zap.Logger
		DB         *DB `keyval:"db,omitempty"`
	}
	Server struct {
		router chi.Router
		conf   Config
		coffee coffee.Server
	}
)

func (s *Server) GetRouter() chi.Router {
	return s.router
}

// routeMap is a mapping that provides a HTTP handler for a given tuple of
// (path, method).
// If "*" is given for method, the handler will match for any HTTP method for
// the given path.
var routeMap = types.RouteMap[*Server]{
	{"/api/coffee", "*"}: handleCoffee,
	{"/api", "GET"}:      handleAPIIndex,
	{"/*", "GET"}:        handleStatic,
}

func NewServer(cfg Config) *Server {
	var coffeeStore coffee.Store
	if cfg.DB.Enabled {
		coffeeStore = coffee.NewDBStore(cfg.DB.Conn)
	} else {
		coffeeStore = coffee.NewMemoryStore()
	}
	s := &Server{
		router: chi.NewRouter(),
		conf:   cfg,
		coffee: *coffee.New(
			coffeeStore,
			cfg.Logger.With(zap.String("component", "coffee")),
		),
	}

	// Middleware
	s.router.Use(middleware.NewLoggerMiddleware(s.conf.Logger))
	routeMap.Build(s)

	return s
}

func (s *Server) Serve() {
	addr := net.JoinHostPort(s.conf.Host, strconv.Itoa(s.conf.Port))
	s.conf.Logger.Info("serving http",
		zap.String("hostname", s.conf.Host),
		zap.Int("port", s.conf.Port),
	)
	http.ListenAndServe(addr, s.router)
}

func handleAPIIndex(_ *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("{}"))
	}
}

func handleStatic(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.EscapedPath()[1:]
		if path == "" {
			path = "index.html"
		}
		filePath := s.conf.StaticPath + "/" + path
		s.conf.Logger.Info("trying path for static file",
			zap.String("path", filePath),
		)
		http.ServeFile(w, r, filePath)
	}
}

func handleCoffee(s *Server) http.HandlerFunc {
	return s.coffee.ServeHTTP
}
