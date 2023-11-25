package server

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hiMaisie/maisie-site/internal/pkg/middleware"
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

var routeMap = map[route]func(*Server) http.HandlerFunc{
	{"/", "GET"}:  handleIndex,
	{"/*", "GET"}: handleStatic,
}

func NewServer(cfg Config) *Server {
	s := &Server{
		router: chi.NewRouter(),
		conf:   cfg,
	}

	// Middleware
	s.router.Use(middleware.NewLoggerMiddleware(s.conf.Logger))

	for r, handler := range routeMap {
		s.router.Method(r.method, r.path, handler(s))
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

func handleIndex(_ *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("{}"))
	}
}

func handleStatic(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := s.conf.StaticPath + "/" + r.URL.EscapedPath()
		s.conf.Logger.Info("trying path for static file",
			zap.String("path", path),
		)
		http.ServeFile(w, r, path)
	}
}
