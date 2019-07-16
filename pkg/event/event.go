package event

import "time"

// Event -
type Event interface {
	EventID() string
	AggregateID() string
	Version() int
	CreatedAt() time.Time
	Source() Source
}

// Source -
type Source map[string]string
