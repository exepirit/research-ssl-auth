package authn

import "net/http"

// Authenticator аутентифицирует пользователя
type Authenticator interface {
	// AuthHTTP аутентифицирует пользователя по данным в HTTP-запросе.
	AuthHTTP(r *http.Request) (User, error)
}
