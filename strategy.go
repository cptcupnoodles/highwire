package main

import (
	"fmt"
	"sync"
)

type Strategy struct {
	mu   sync.Mutex
	next int
}

func (r *Strategy) RoundRobin(backends []*Backend) (*Backend, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(backends) == 0 {
		return nil, fmt.Errorf("no backends configured")
	}

	for checked := 0; checked < len(backends); checked++ {
		index := r.next
		r.next = (r.next + 1) % len(backends)

		if backends[index].Health() {
			return backends[index], nil
		}
	}

	return nil, fmt.Errorf("no healthy backends available")
}

func roundWeightedRobin(backends []Backend) {}
