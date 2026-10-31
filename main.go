package main

import (
	"log"
	"net/http"
)

func main() {

	var (
		backends = []*Backend{
			AddBackend("alpha", "http://127.0.0.1:9000"),
			AddBackend("beta", "http://127.0.0.1:9001"),
			AddBackend("charlie", "http://127.0.0.1:9002"),
		}
		//nextBackend int
	)

	// Health check all the registered backends
	for i := range backends {
		isHealthy := backends[i].healthCheck()
		if isHealthy {
			backends[i].MarkHealthy()
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Implement a backend seletion strategy
		backend := backends[0] // temporary

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
