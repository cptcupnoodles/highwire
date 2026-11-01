package main

import (
	"fmt"
	"sync"
)

type Backend struct {
	mu sync.Mutex

	Name                string
	Address             string
	IsHealthy           bool
	ActiveRequest       int
	TotalStartedRequest int
	Failure             int
	Weight              int
}

func AddBackend(name, address string, weight int) *Backend {
	backend := &Backend{
		Name:                name,
		Address:             address,
		IsHealthy:           false,
		ActiveRequest:       0,
		TotalStartedRequest: 0,
		Failure:             0,
		Weight:              weight,
	}

	return backend
}

func (b *Backend) RemoveBackend() error {
	return fmt.Errorf("Cannot remove active backend")
}

func (b *Backend) Health() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.IsHealthy
}

func (b *Backend) MarkHealthy() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.IsHealthy = true
}

func (b *Backend) MarkUnhealthy() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.IsHealthy = false
}

func (b *Backend) Start() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.IsHealthy {
		return fmt.Errorf("backend %s is unhealthy", b.Name)
	}

	b.ActiveRequest++
	b.TotalStartedRequest++

	return nil
}

func (b *Backend) Complete() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.ActiveRequest == 0 {
		return fmt.Errorf("no active request for the backend %s ", b.Name)
	}

	b.ActiveRequest--
	return nil
}

func (b *Backend) Fail() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.ActiveRequest == 0 {
		return fmt.Errorf("no active request for the backend %s ", b.Name)
	}
	b.ActiveRequest--
	b.Failure++

	return nil
}

func TotalWeight(backends []*Backend) (int, error) {
	var totalWeight int

	if len(backends) == 0 {
		return 0, fmt.Errorf("no backend configured")
	}

	for _, value := range backends {
		totalWeight += value.Weight
	}

	return totalWeight, nil
}

func (b *Backend) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return fmt.Sprintf("name=%s name=%s: %v, %d, %d, %d",
		b.Name,
		b.Address,
		b.IsHealthy,
		b.ActiveRequest,
		b.TotalStartedRequest,
		b.Failure,
	)
}
