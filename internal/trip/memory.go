package trip

import (
	"cmp"
	"slices"
	"sync"
)

type MemoryStore struct {
	mu     sync.RWMutex
	trips  map[int]Trip
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{trips: make(map[int]Trip), nextID: 1}
}

func (s *MemoryStore) List() []Trip {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Trip, 0, len(s.trips))
	for _, t := range s.trips {
		out = append(out, t)
	}
	slices.SortFunc(out, func(a, b Trip) int { return cmp.Compare(a.ID, b.ID) })
	return out
}

func (s *MemoryStore) Get(id int) (Trip, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.trips[id]
	if !ok {
		return Trip{}, ErrNotFound
	}
	return t, nil
}

func (s *MemoryStore) Create(t Trip) Trip {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ID = s.nextID
	s.nextID++
	s.trips[t.ID] = t
	return t
}

func (s *MemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.trips[id]; !ok {
		return ErrNotFound
	}
	delete(s.trips, id)
	return nil
}
