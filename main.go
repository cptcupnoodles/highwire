package main

import "net/http"

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
		backend := backends[0]

		if err := backend.Start(); err != nil {

		}

	})

}
