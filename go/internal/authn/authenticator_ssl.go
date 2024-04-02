package authn

import (
	"crypto/x509"
	"errors"
	"log/slog"
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
	logger := slog.With("module", "authn")

	if r.TLS == nil {
		logger.Debug("request has no TLS session")
		return auth.createUnauthenticatedUser(ErrNoTLS)
	}

	certs := r.TLS.PeerCertificates
	for _, cert := range certs {
		logger.Debug("found valid client certificate", "subject", cert.Subject.String())
		return auth.createUserFromCert(cert), nil
	}

	logger.Debug("request has no valid client certificate")
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
