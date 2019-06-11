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
	http.ListenAndServe(":8081", http.DefaultServeMux)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	valStr := r.URL.Query().Get("vals")
	log.Printf("got values: %s \n", valStr)

	vals := strings.Split(valStr, ",")

	ints := make([]int, 0, len(vals))
	for _, val := range vals {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("can't parse val: %s, skipping \n", val)
			continue
		}

		ints = append(ints, valInt)
	}

	acc := 0.0
	for _, val := range ints {
		acc = acc + float64(val)
	}

	w.Write([]byte("results: \n"))
	fmt.Fprintf(w, "sum of %v is %f\n", valStr, acc)
	fmt.Fprintf(w, "avg of %v is %f\n", valStr, acc/float64(len(vals)))
}
