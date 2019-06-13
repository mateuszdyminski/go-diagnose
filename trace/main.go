package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"
)

const hostPort = ":8090"

func main() {
	flag.Parse()

	http.HandleFunc("/hello", hello)

	log.Printf("Starting HTTP server on port %s \n", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

var mutex = &sync.Mutex{}
var counter = 0

func hello(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	counter = counter + 1

	// some heavy operation goes here
	time.Sleep(10 * time.Millisecond)

	fmt.Fprintf(w, "Hello world. Endpoint visited %d times!", counter)
}
