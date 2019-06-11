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
	log.Printf("values after split: %v \n", vals)

	acc := 0.0
	for i, val := range vals {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("can't parse val: %s err: %s, skipping \n", val, err)
			continue
		} else {
			log.Printf("val: %s parsed properly\n", val)
		}

		acc = acc + float64(valInt)
		log.Printf("temp sum after %d of values is: %f\n", i+1, acc)
	}

	log.Printf("sum of %s is: %v \n", valStr, vals)

	w.Write([]byte("results: \n"))
	fmt.Fprintf(w, "sum of %v is %f\n", valStr, acc)
}
