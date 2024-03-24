package middleware

import (
	"context"
	"github.com/exepirit/research-ssl-auth/go/internal/authn"
	"net/http"
)

// NewAuthMiddleware создает новую аутентифицирующую middleware.
func NewAuthMiddleware(handler http.Handler, authenticator authn.Authenticator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := authenticator.AuthHTTP(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}
