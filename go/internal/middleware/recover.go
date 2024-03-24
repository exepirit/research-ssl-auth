package middleware

import (
	"log/slog"
	"net/http"
)

// NewRecoverMiddleware создает новую middleware, обрабатывающую паники.
func NewRecoverMiddleware(handler http.Handler) http.Handler {
	logger := slog.With("module", "http.server")

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				logger.Error("Panic occurred in request handler", "error", err)
			}
		}()

		handler.ServeHTTP(writer, request)
	})
}
