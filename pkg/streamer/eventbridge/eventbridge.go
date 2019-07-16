package eventbridge

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	awseventbridge "github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

// EventBridgeStreamer -
type EventBridgeStreamer struct {
	stream *awseventbridge.EventBridge
	source string
	ebName string
}

func NewEventBridgeStreamer(opts ...EventBridgeStreamerOption) (streamer.Streamer, error) {

	result := EventBridgeStreamer{}
	for _, opt := range opts {
		if err := opt(&result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (streamer EventBridgeStreamer) Push(ctx context.Context, records []streamer.Record) error {

	ebEntries := []*awseventbridge.PutEventsRequestEntry{}

	for _, record := range records {
		ebEntries = append(ebEntries, &awseventbridge.PutEventsRequestEntry{
			Detail:       aws.String(string(record.Data)),
			DetailType:   aws.String(record.Type),
			EventBusName: aws.String(streamer.ebName),
			Source:       aws.String(streamer.source),
			Time:         aws.Time(time.Now().UTC()),
		})
	}

	input := &awseventbridge.PutEventsInput{}

	output, err := streamer.stream.PutEventsWithContext(ctx, input)
	if err != nil {
		return err
	}

	if *output.FailedEntryCount != 0 {
		return errors.New("unable to write to eventbridge")
	}

	return nil
}
