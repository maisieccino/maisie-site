package types

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type fakeServer struct {
	chi.Router
	getHandlerCalled      bool
	wildcardHandlerCalled bool
}

func (s *fakeServer) GetRouter() chi.Router {
	return s
}

func Test_RouteMapBuildRoutes(t *testing.T) {
	r := fakeServer{
		chi.NewRouter(),
		false,
		false,
	}

	mockGETHandler := func(s *fakeServer) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			s.getHandlerCalled = true
			w.WriteHeader(http.StatusOK)
		}
	}

	mockWildcardHandler := func(s *fakeServer) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			s.wildcardHandlerCalled = true
			w.WriteHeader(http.StatusOK)
		}
	}

	routeMap := RouteMap[*fakeServer]{
		{Method: "GET", Path: "/test"}: mockGETHandler,
		{Method: "*", Path: "/test2"}:  mockWildcardHandler,
	}

	routeMap.Build(&r)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.True(t, r.getHandlerCalled)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPatch, "/test2", nil))
	assert.True(t, r.wildcardHandlerCalled)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
}
