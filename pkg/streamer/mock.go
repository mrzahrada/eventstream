package streamer

import (
	"context"
	"errors"
)

type MockStreamer struct{}

func NewMockStreamer() (Streamer, error) {
	return MockStreamer{}, nil
}

func (streamer MockStreamer) Push(ctx context.Context, records []Record) error {
	return errors.New("not implemented")
}
