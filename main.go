package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	port, err := readPortFromEnv("HTTP_ECH0_PORT", 80)
	if err != nil {
		log.Fatal("Failed to parse HTTP_ECH0_PORT environment variable")
	}
	tlsPort, err := readPortFromEnv("HTTP_ECH0_TLS_PORT", 443)
	if err != nil {
		log.Fatalf("Failed to parse HTTP_ECH0_TLS_PORT environment variable")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	http.HandleFunc("/", index)
	log.Printf("Running on port %d and %d (TLS)\n", port, tlsPort)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Printf("Failed to start server on port %d", port)
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		err := http.ListenAndServeTLS(fmt.Sprintf(":%d", tlsPort), "server.crt", "server.key", nil)
		if err != nil {
			log.Printf("Failed to start server on port %d", tlsPort)
			log.Fatal(err)
		}
	}()
	wg.Wait()
}

func readPortFromEnv(name string, def int) (int, error) {
	value, set := os.LookupEnv(name)
	if !set {
		return def, nil
	}
	return strconv.Atoi(value)
}

func index(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
	io.Copy(w, r.Body)
}
