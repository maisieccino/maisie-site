package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// NewLoggerMiddleware adds a log line to requests that the server handles.
func NewLoggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				end := time.Now()
				duration := end.Sub(start)
				logger.Info("handled request",
					zap.String("method", r.Method),
					zap.String("status", w.Header().Get("Status")),
					zap.Int("duration_ms", int(duration.Milliseconds())),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
