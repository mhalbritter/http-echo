package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Printf("Running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
	io.Copy(w, r.Body)
}
