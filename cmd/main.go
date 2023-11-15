package main

import (
	"flag"

	"github.com/superjcd/calendarservice/cmd/server"
)

var cfg = flag.String("config", "conf/config.yaml", "config file location")

func main() {
	flag.Parse()
	server.Run(*cfg)
}
