package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/add", http.HandlerFunc(addHandler))

	log.Println("Starting HTTP server on port 8081")
	if err := http.ListenAndServe(":8081", http.DefaultServeMux); err != nil {
		log.Fatalln(err.Error())
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	valStr := r.URL.Query().Get("vals")
	vals := strings.Split(valStr, ",")

	acc := 0.0
	for _, val := range vals {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			continue
		}

		acc = acc + float64(valInt)
	}

	w.Write([]byte("results: \n"))
	fmt.Fprintf(w, "sum of %v is %f\n", valStr, acc)
}
