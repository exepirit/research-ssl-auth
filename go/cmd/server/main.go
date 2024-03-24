package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/exepirit/research-ssl-auth/go/internal/authn"
	"github.com/exepirit/research-ssl-auth/go/internal/middleware"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	))

	cfg, err := loadFlagsConfig()
	if err != nil {
		slog.Error("Load config error", "error", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/hello", HelloHandler{})

	handler := middleware.NewAuthMiddleware(mux, authn.SSLAuthenticator{
		RootCertificates: &x509.CertPool{},
		AllowAnonymous:   true,
	})
	handler = middleware.NewRecoverMiddleware(handler)
	handler = middleware.NewLoggingMiddleware(handler)

	server := http.Server{
		Addr:    cfg.ListenAddr,
		Handler: handler,
		TLSConfig: &tls.Config{
			ServerName: cfg.Host,
			ClientAuth: tls.RequestClientCert, // request, but not require client certificate
		},
	}
	if err := server.ListenAndServeTLS(cfg.CertPath, cfg.CertKeyPath); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}

func loadCertificates(filesystem fs.FS, filenames []string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()
	for _, fname := range filenames {
		file, err := filesystem.Open(fname)
		if err != nil {
			return pool, fmt.Errorf("failed to open file %q: %w", fname, err)
		}

		content, err := io.ReadAll(file)
		if err != nil {
			return pool, fmt.Errorf("failed to load certificate content from %q: %w", fname, err)
		}

		block, _ := pem.Decode(content)
		if block == nil {
			return pool, fmt.Errorf("file %q does not contain any PEM data", fname)
		}

		certificates, err := x509.ParseCertificates(block.Bytes)
		if err != nil {
			return pool, fmt.Errorf("x509 certificate parsing error: %w", err)
		}

		for _, cert := range certificates {
			pool.AddCert(cert)
		}
	}
	return pool, nil
}
