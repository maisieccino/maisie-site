package server

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maisieccino/maisie-site/internal/pkg/middleware"
	"go.uber.org/zap"
)

type (
	Config struct {
		Host       string `keyval:"host"`
		Port       int    `keyval:"port"`
		StaticPath string `keyval:"staticPath"`
		Logger     *zap.Logger
	}
	Server struct {
		router chi.Router
		conf   Config
	}
	route struct {
		path   string
		method string
	}
)

// routeMap is a mapping that provides a HTTP handler for a given tuple of
// (path, method).
// If "*" is given for method, the handler will match for any HTTP method for
// the given path.
var routeMap = map[route]func(*Server) http.HandlerFunc{
	{"/api/coffee", "*"}: handleCoffee,
	{"/api", "GET"}:      handleAPIIndex,
	{"/*", "GET"}:        handleStatic,
}

func NewServer(cfg Config) *Server {
	s := &Server{
		router: chi.NewRouter(),
		conf:   cfg,
	}

	// Middleware
	s.router.Use(middleware.NewLoggerMiddleware(s.conf.Logger))

	for r, handler := range routeMap {
		if r.method == "*" {
			s.router.Mount(r.path, handler(s))
		} else {
			s.router.Method(r.method, r.path, handler(s))
		}
	}

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
	// TODO: Add in the coffee map router
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
