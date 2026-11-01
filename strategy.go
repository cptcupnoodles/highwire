package main

import (
	"fmt"
	"sync"
)

type Strategy struct {
	mu   sync.Mutex
	next int
}

func (s *Strategy) RoundRobin(backends []*Backend) (*Backend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(backends) == 0 {
		return nil, fmt.Errorf("no backends configured")
	}

	for range len(backends) {
		index := s.next
		s.next = (s.next + 1) % len(backends)

		if backends[index].Health() {
			return backends[index], nil
		}
	}

	return nil, fmt.Errorf("no healthy backends available")
}

func (s *Strategy) WeightedRoundRobin(backends []*Backend) (*Backend, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(backends) == 0 {
		return nil, fmt.Errorf("no backends configured")
	}
}
