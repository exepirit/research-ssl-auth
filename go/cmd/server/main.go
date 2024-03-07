package main

import (
	"crypto/tls"
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

	handler := recoverMiddleware(mux)
	handler = loggingMiddleware(mux)

	server := http.Server{
		Addr:    cfg.ListenAddr,
		Handler: handler,
		TLSConfig: &tls.Config{
			ServerName: cfg.Host,
		},
	}
	if err := server.ListenAndServeTLS(cfg.CertPath, cfg.CertKeyPath); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}
