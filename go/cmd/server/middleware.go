package main

import (
	"log/slog"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(handler http.Handler) http.Handler {
	logger := slog.With("module", "http.server")

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()

		loggingWriter := &loggingResponseWriter{ResponseWriter: writer}
		handler.ServeHTTP(loggingWriter, request)

		path := request.URL.Path
		latency := time.Now().Sub(start)
		responseStatus := loggingWriter.statusCode

		logger.Info("Request handled",
			slog.String("path", path),
			slog.Duration("latency", latency),
			slog.Int("statusCode", responseStatus),
		)
	})
}

func recoverMiddleware(handler http.Handler) http.Handler {
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
