package authn

import (
	"crypto/x509"
	"errors"
	"net/http"
	"time"
)

var (
	ErrNoTLS              = errors.New("request has no TLS session")
	ErrNoValidCertificate = errors.New("session has no valid certificate")
)

// SSLAuthenticator реализует аутентификацию используя SSL/TLS-сертификаты.
type SSLAuthenticator struct {
	RootCertificates *x509.CertPool
	AllowAnonymous   bool
}

func (auth SSLAuthenticator) AuthHTTP(r *http.Request) (User, error) {
	if r.TLS == nil {
		return auth.createUnauthenticatedUser(ErrNoTLS)
	}

	certs := r.TLS.PeerCertificates
	for _, cert := range certs {
		if auth.validateCertificate(cert) {
			return auth.createUserFromCert(cert), nil
		}
	}

	return auth.createUnauthenticatedUser(ErrNoValidCertificate)
}

func (auth SSLAuthenticator) validateCertificate(crt *x509.Certificate) bool {
	_, err := crt.Verify(x509.VerifyOptions{
		Roots:       auth.RootCertificates,
		CurrentTime: time.Now(),
	})
	if err != nil {
		return false
	}

	return true
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
