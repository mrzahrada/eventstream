package streamer

import "context"

// Streamer -
type Streamer interface {
	Push(ctx context.Context, records []Record) error
}
