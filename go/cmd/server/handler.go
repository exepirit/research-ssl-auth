package main

import (
	"encoding/json"
	"github.com/exepirit/research-ssl-auth/go/internal/authn"
	"github.com/exepirit/research-ssl-auth/go/internal/middleware"
	"net/http"
)

type TLSStateDto struct {
	CertificatesSubjects []string `json:"certificatesSubjects"`
}

type UserDto struct {
	ID string `json:"id"`
}

type AboutMeResponse struct {
	User *UserDto     `json:"user"`
	TLS  *TLSStateDto `json:"tls"`
}

// AboutMeHandler обрабатывает запрос и возвращает информацию о пользователе и соединении в сериализованой
// модели  AboutMeResponse.
func AboutMeHandler(writer http.ResponseWriter, request *http.Request) {
	response := AboutMeResponse{}

	user, authenticated := middleware.GetCurrentUser(request.Context())
	if authenticated {
		switch user.Type() {
		case authn.UserTypeAnonymous:
			response.User = &UserDto{ID: "anonymous"}
		case authn.UserTypeAuthenticated:
			response.User = &UserDto{ID: user.ID()}
		default:
			panic("unexpected user type")
		}
	}

	if request.TLS != nil {
		response.TLS = &TLSStateDto{}

		response.TLS.CertificatesSubjects = make([]string, len(request.TLS.PeerCertificates))
		for i, cert := range request.TLS.PeerCertificates {
			response.TLS.CertificatesSubjects[i] = cert.Subject.String()
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}
