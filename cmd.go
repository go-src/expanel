package main

import (
	"flag"
	"fmt"
)

func parseFlags(cfg *Config) {
	fmt.Println(cfg)
	flag.StringVar(&cfg.Http.Addr, "addr", cfg.Http.Addr, "Bind address for web server")
	flag.StringVar(&cfg.Http.Path, "path", cfg.Http.Path, "Static root for web server")
	flag.BoolVar(&cfg.Http.TLS, "tls", cfg.Http.TLS, "Enable TLS support for web server")
	flag.BoolVar(&cfg.Runtime.Daemon, "daemon", cfg.Runtime.Daemon, "Run web server as a daemon")
	flag.IntVar(&cfg.Runtime.Procs, "procs", cfg.Runtime.Procs, "runtime.GOMAXPROCS(x)")
	flag.Parse()
	fmt.Println(cfg)
}
