package trip

import "errors"

var ErrNotFound = errors.New("trip not found")

type Trip struct {
	ID          int    `json:"id"`
	Destination string `json:"destination"`
	Notes       string `json:"notes,omitempty"`
}

// Store is an interface so we can swap in Postgres later.
type Store interface {
	List() []Trip
	Get(id int) (Trip, error)
	Create(t Trip) Trip
	Delete(id int) error
}
