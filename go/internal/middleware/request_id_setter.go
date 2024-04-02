package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const RequestIDKey = "requestId"

// NewSetRequestIDMiddleware создает middleware устанавливающую ID запроса.
func NewSetRequestIDMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id := uuid.New() // TODO: use 128-bit counter instead pseudo-random UUIDv4
		ctx := context.WithValue(request.Context(), RequestIDKey, id.String())
		request = request.WithContext(ctx)
		handler.ServeHTTP(writer, request)
	})
}

// GetRequestID возвращает ID запроса или пустую строку, если он не установлен.
func GetRequestID(ctx context.Context) string {
	value := ctx.Value(RequestIDKey)
	if value == nil {
		return ""
	}
	requestID, ok := value.(string)
	if !ok {
		return ""
	}
	return requestID
}
