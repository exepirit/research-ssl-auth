package main

import "flag"

// Config описывает конфигурацию приложения.
type Config struct {
	CertPath       string
	CertKeyPath    string
	ListenAddr     string
	Host           string
	AllowAnonymous bool
}

func loadFlagsConfig() (Config, error) {
	var cfg Config
	flag.StringVar(&cfg.CertPath, "cert", "certificate.crt", "Server certificate file path")
	flag.StringVar(&cfg.CertKeyPath, "key", "certificate.key", "Server certificate key file path")
	flag.StringVar(&cfg.ListenAddr, "listen", "localhost:8080", "Server listen address")
	flag.StringVar(&cfg.Host, "host", "localhost", "Server host")
	flag.BoolVar(&cfg.AllowAnonymous, "anonymous", false, "Allow anonymous users")
	flag.Parse()

	return cfg, nil
}
