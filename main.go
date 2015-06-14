package main

import (
	"fmt"
)

func main() {
	cfg := getConfig()
	parseFlags(&cfg)
	fmt.Println(cfg)
	//startServer(cfg.Http.Addr, cfg.Http.Path)
}
