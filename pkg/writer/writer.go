package writer

import (
	"context"

	"github.com/mrzahrada/eventstream/pkg/event"
	"github.com/mrzahrada/eventstream/pkg/serializer"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

// Writer -
type Writer struct {
	streamer   streamer.Streamer
	serializer serializer.Serializer

	buffer []event.Event
}

// Commit event to writer
func (writer *Writer) Commit(event event.Event) {
	writer.buffer = append(writer.buffer, event)
}

// Push data to streamer
func (writer Writer) Push(ctx context.Context) error {

	defer writer.Clear()

	records := []streamer.Record{}

	for _, event := range writer.buffer {
		record, err := writer.serializer.Marshal(event)
		if err != nil {
			return err
		}
		records = append(records, record)
	}

	return writer.streamer.Push(ctx, records)
}

// Clear buffer
func (writer Writer) Clear() {
	writer.buffer = writer.buffer[:0]
}

// Len returns number of events in buffer
func (writer Writer) Len() int {
	return len(writer.buffer)
}
