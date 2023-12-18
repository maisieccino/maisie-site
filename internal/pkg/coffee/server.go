package coffee

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

const storeTimeout = time.Second * 2

func writeError(w http.ResponseWriter, logger *zap.Logger, err error, statusCode int) {
	w.WriteHeader(statusCode)
	errorObj := struct{ Err string }{Err: err.Error()}
	enc := json.NewEncoder(w)
	if encErr := enc.Encode(errorObj); encErr != nil {
		logger.Error("error writing error response body", zap.Error(encErr))
	}
}

type Server struct {
	store Store
	chi.Router
	logger *zap.Logger
}

type route struct {
	path   string
	method string
}

var routeMap = map[route]func(*Server) http.HandlerFunc{
	{"/places", "GET"}: handleListPlaces,
}

func handleListPlaces(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		innerCtx, cancel := context.WithTimeout(r.Context(), storeTimeout)
		defer cancel()
		items, err := s.store.List(innerCtx)
		if err != nil {
			s.logger.Error("error listing map items", zap.Error(err))
			writeError(w, s.logger, err, http.StatusInternalServerError)
		}

		resp := struct{ Items []MapItem }{Items: items}
		enc := json.NewEncoder(w)
		if encErr := enc.Encode(resp); encErr != nil {
			s.logger.Error("error writing body", zap.Error(err))
			writeError(w, s.logger, err, http.StatusInternalServerError)
			return
		}
	}
}

func New(store Store, logger *zap.Logger) *Server {
	router := chi.NewRouter()
	return &Server{
		store:  store,
		Router: router,
		logger: logger.With(zap.String("component", "coffee")),
	}
}

func (s *Server) GetRouter() chi.Router {
	return s.Router
}
