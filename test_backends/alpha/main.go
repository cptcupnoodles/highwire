package main

import (
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from alpha"))
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}
