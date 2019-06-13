package main

import (
	"flag"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/mateuszdyminski/go-diagnose/profile/handlers"
)

const hostPort = ":8090"

func main() {
	flag.Parse()

	http.HandleFunc("/statsHello", handlers.WithStats(handlers.Hello))
	http.HandleFunc("/hello", handlers.Hello)

	log.Printf("Starting HTTP server on port %s \n", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
