package main

import "flag"

type Config struct {
	CertPath    string
	CertKeyPath string
	ListenAddr  string
	Host        string
}

func loadFlagsConfig() (Config, error) {
	var cfg Config
	flag.StringVar(&cfg.CertPath, "cert", "certificate.crt", "Certificate file path")
	flag.StringVar(&cfg.CertKeyPath, "key", "certificate.key", "Certificate key file path")
	flag.StringVar(&cfg.ListenAddr, "listen", "localhost:8080", "Server listen address")
	flag.StringVar(&cfg.Host, "host", "localhost", "Server host")
	flag.Parse()

	return cfg, nil
}
