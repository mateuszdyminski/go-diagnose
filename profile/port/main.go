package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/mateuszdyminski/go-diagnose/profile/handlers"
)

const hostPort = ":8090"

func main() {
	flag.Parse()

	// pprof configuration
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", pprof.Index)
		mux.HandleFunc("/cmdline", pprof.Cmdline)
		mux.HandleFunc("/profile", pprof.Profile)
		mux.HandleFunc("/symbol", pprof.Symbol)
		mux.HandleFunc("/trace", pprof.Trace)
		log.Println("Starting profile HTTP server on port 1337")
		if err := http.ListenAndServe(":1337", mux); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// register handlers
	http.HandleFunc("/statsHello", handlers.WithStats(handlers.Hello))
	http.HandleFunc("/hello", handlers.Hello)

	log.Printf("Starting HTTP server on port %s \n", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
