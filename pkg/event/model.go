package event

import (
	"time"
)

// Model -
type Model struct {
	EID string    `json:"_eid"`
	AID string    `json:"_aid"`
	At  time.Time `json:"_at"`
	Src Source    `json:"_s"`
	Ver int       `json:"_v"`
}

// EventID -
func (model Model) EventID() string {
	return model.EID
}

// AggregateID -
func (model Model) AggregateID() string {
	return model.AID
}

// Version -
func (model Model) Version() int {
	return model.Ver
}

// CreatedAt -
func (model Model) CreatedAt() time.Time {
	return model.At
}

// Source -
func (model Model) Source() Source {
	return model.Src
}
