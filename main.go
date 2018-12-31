package main

import (
	"fmt"
	"net/http"
)

func getHello() string {
	return "hello world!!"
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World")
	fmt.Fprintf(w, "hello world")
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
