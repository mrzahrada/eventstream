package serializer

import (
	"github.com/mrzahrada/eventstream/pkg/event"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

// Serializer -
type Serializer interface {
	Marshal(event.Event) (streamer.Record, error)
	Unmarshal(streamer.Record) (event.Event, error)
}
