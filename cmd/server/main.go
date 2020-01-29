package main

import (
	"flag"
	"log"

	"github.com/andreylm/mind_systems/pkg/server"
)

var (
	port string
	host string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Host")
	flag.StringVar(&port, "port", "8000", "Port")
	flag.Parse()
}

func main() {
	srv := server.NewServer(host, port)
	log.Panic(srv.Start())
}
