package authn

import (
	"crypto/x509"
	"errors"
	"net/http"
)

var (
	ErrNoTLS              = errors.New("request has no TLS session")
	ErrNoValidCertificate = errors.New("session has no valid certificate")
)

// SSLAuthenticator реализует аутентификацию используя SSL/TLS-сертификаты.
type SSLAuthenticator struct {
	AllowAnonymous bool
}

func (auth SSLAuthenticator) AuthHTTP(r *http.Request) (User, error) {
	if r.TLS == nil {
		return auth.createUnauthenticatedUser(ErrNoTLS)
	}

	certs := r.TLS.PeerCertificates
	for _, cert := range certs {
		return auth.createUserFromCert(cert), nil
	}

	return auth.createUnauthenticatedUser(ErrNoValidCertificate)
}

func (SSLAuthenticator) createUserFromCert(crt *x509.Certificate) User {
	return AuthenticatedUser{
		id: crt.Subject.String(),
	}
}

func (auth SSLAuthenticator) createUnauthenticatedUser(reason error) (User, error) {
	if auth.AllowAnonymous {
		return AnonymousUser{}, nil
	}
	return nil, reason
}
