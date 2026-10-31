package main

import (
	"log"
	"net/http"
)

func main() {

	backends := []*Backend{
		AddBackend("alpha", "http://127.0.0.1:9000", 2),
		AddBackend("beta", "http://127.0.0.1:9001", 3),
		AddBackend("charlie", "http://127.0.0.1:9002", 5),
	}

	strategy := &Strategy{}

	// Health check all the registered backends
	for i := range backends {
		isHealthy := backends[i].healthCheck()
		if isHealthy {
			backends[i].MarkHealthy()
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Implement a backend seletion strategy
		backend, err := strategy.RoundRobin(backends)

		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		if err := backend.Start(); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		defer func() {
			_ = backend.Complete()
		}()

		// Implement the proxy
		backend.proxy(w, r)

	})

	log.Println("proxy listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
