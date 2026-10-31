package strategy

import "sync"

type Backend struct {
	mu sync.Mutex

	Name                string
	Address             string
	IsHealthy           bool
	ActiveRequest       int
	TotalStartedRequest int
	Failure             int
}

func RoundRobin(backends []Backend) {

}
