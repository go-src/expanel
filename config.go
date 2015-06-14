package main

import (
	"log"

	"github.com/go-src/gcfg"
)

type Config struct {
	Http struct {
		Addr string
		Path string
		TLS  bool
	}
	Runtime struct {
		Daemon bool
		Procs  int
	}
}

func getConfig() Config {
	cfg := Config{}
	err := gcfg.ReadFileInto(&cfg, "expanel.conf")
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	return cfg
}
