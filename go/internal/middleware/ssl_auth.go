package middleware

import (
	"context"
	"github.com/exepirit/research-ssl-auth/go/internal/authn"
	"net/http"
)

const ctxUserKey = "user"

// NewAuthMiddleware создает новую аутентифицирующую middleware.
func NewAuthMiddleware(handler http.Handler, authenticator authn.Authenticator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := authenticator.AuthHTTP(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserKey, user)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

// GetCurrentUser извлекает текущего пользователя, указанного в контексте.
func GetCurrentUser(ctx context.Context) (authn.User, bool) {
	value := ctx.Value(ctxUserKey)
	if value == nil {
		return nil, false
	}
	user, ok := value.(authn.User)
	return user, ok
}
