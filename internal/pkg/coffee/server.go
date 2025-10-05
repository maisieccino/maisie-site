package coffee

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/maisieccino/maisie-site/internal/pkg/types"
	"github.com/maisieccino/maisie-site/pkg/api"
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

func (s *Server) GetRouter() chi.Router {
	return s.Router
}

// routeMap is a mapping that provides a HTTP handler for a given tuple of
// (path, method).
// If "*" is given for method, the handler will match for any HTTP method for
// the given path.
var routeMap = types.RouteMap[*Server]{
	{Path: "/places", Method: "GET"}:         handleListPlaces,
	{Path: "/places", Method: "PUT"}:         handleCreatePlace,
	{Path: "/places/:id", Method: "PATCH"}:   handleUpdatePlace,
	{Path: "/places/by-area", Method: "PUT"}: handleSearchByArea,
}

// TODO: Move to a marshalling package
func itemToJSON(i MapItem) api.Place {
	return api.Place{
		Id:        uuid.MustParse(i.ID),
		ImageUrl:  &i.ImageURL,
		Latitude:  float32(i.Latitude),
		Longitude: float32(i.Longitude),
		Name:      i.Name,
		ReviewUrl: &i.ReviewURL,
		Type:      string(i.Type),
	}
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

		places := []api.Place{}
		for _, i := range items {
			places = append(places, itemToJSON(i))
		}

		resp := struct {
			Items []api.Place `json:"items"`
		}{Items: places}
		enc := json.NewEncoder(w)
		if encErr := enc.Encode(resp); encErr != nil {
			s.logger.Error("error writing body", zap.Error(err))
			writeError(w, s.logger, err, http.StatusInternalServerError)
			return
		}
	}
}

func handleCreatePlace(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, s.logger, err, http.StatusBadRequest)
			return
		}
		var place MapItem
		if err := json.Unmarshal(body, &place); err != nil {
			writeError(w, s.logger, fmt.Errorf("unmarshaling place: %w", err), http.StatusBadRequest)
			return
		}

		if err := s.store.Create(r.Context(), place); err != nil {
			writeError(w, s.logger, fmt.Errorf("storing place: %w", err), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"success"}`))
	}
}

func handleUpdatePlace(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, s.logger, err, http.StatusBadRequest)
			return
		}

		var newItem MapItem
		if err := json.Unmarshal(body, &newItem); err != nil {
			writeError(w, s.logger, err, http.StatusBadRequest)
			return
		}

		if err := s.store.Update(r.Context(), id, newItem); err != nil {
			writeError(w, s.logger, fmt.Errorf("updating place: %w", err), http.StatusInternalServerError)
		}
	}
}

func handleSearchByArea(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, s.logger, err, http.StatusBadRequest)
			return
		}
		var params SearchByAreaParams
		if err := json.Unmarshal(body, &params); err != nil {
			writeError(w, s.logger, err, http.StatusBadRequest)
			return
		}

		results, err := s.store.SearchByArea(r.Context(), params)
		if err != nil {
			writeError(w, s.logger, err, http.StatusInternalServerError)
			return
		}

		resp := struct {
			Items []MapItem `json:"items"`
		}{Items: results}
		enc := json.NewEncoder(w)
		if encErr := enc.Encode(resp); encErr != nil {
			s.logger.Error("error writing body", zap.Error(err))
			writeError(w, s.logger, err, http.StatusInternalServerError)
			return
		}
	}
}

func New(store Store, logger *zap.Logger) *Server {
	s := &Server{
		store:  store,
		Router: chi.NewRouter(),
		logger: logger.With(zap.String("component", "coffee")),
	}

	// Set handlers.
	routeMap.Build(s)
	return s
}
