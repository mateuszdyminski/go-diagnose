package main

import (
	"fmt"
	"net/http"

	"github.com/mateuszdyminski/go-diagnose/instrument/prom"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	// prometheus handler
	mux.Handle("/metrics", promhttp.Handler())

	// our business logic handlers
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/register", register)

	inst := prom.NewInstrument()
	srv := &http.Server{
		Addr:    ":8095",
		Handler: inst.Wrap(mux),
	}

	prometheus.MustRegister(newUserCounter)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Msgf("HTTP server failed: %v", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

var newUserCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: "business",
		Name:      "new_user",
		Help:      "The total number of new users registered in our system.",
	},
	[]string{"name"},
)

func register(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	newUserCounter.WithLabelValues(name).Inc()

	fmt.Fprintln(w, "New user registered in our system!")
}
