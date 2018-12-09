package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloworld)

	server := &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: mux,
	}

	server.ListenAndServe()
}
