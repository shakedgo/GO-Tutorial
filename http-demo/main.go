package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	id := req.PathValue("id")
	fmt.Fprintf(w, "hello %s\n", id)
	// fmt.Fprintf(w, "Hello \n")
	fmt.Println(req.Method + " " + req.RequestURI)
	fmt.Println("Client got hello")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	mux := http.NewServeMux()

	fmt.Println("Starting HTTP server")
	mux.HandleFunc("/hello/{id}", hello) // FIX
	mux.HandleFunc("/headers", headers)

	fmt.Println("Server started at http://localhost:8091")
	http.ListenAndServe(":8091", mux)
}
