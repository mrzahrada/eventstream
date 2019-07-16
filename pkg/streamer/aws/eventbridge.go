package aws

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

// EventBridgeStreamer -
type EventBridgeStreamer struct {
	stream *eventbridge.EventBridge
	source string
	ebName string
}

func NewEventBridgeStreamer() (streamer.Streamer, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	stream := eventbridge.New(sess, &aws.Config{
		Region: aws.String(endpoints.EuWest1RegionID),
	})

	return EventBridgeStreamer{
		stream: stream,
	}, nil

}

func (streamer EventBridgeStreamer) Push(ctx context.Context, records []streamer.Record) error {

	ebEntries := []*eventbridge.PutEventsRequestEntry{}

	for _, record := range records {
		ebEntries = append(ebEntries, &eventbridge.PutEventsRequestEntry{
			Detail:       aws.String(string(record.Data)),
			DetailType:   aws.String(record.Type),
			EventBusName: aws.String(streamer.ebName),
			Source:       aws.String(streamer.source),
			Time:         aws.Time(time.Now().UTC()),
		})
	}

	input := &eventbridge.PutEventsInput{}

	output, err := streamer.stream.PutEventsWithContext(ctx, input)
	if err != nil {
		return err
	}

	if *output.FailedEntryCount != 0 {
		return errors.New("unable to write to eventbridge")
	}

	return nil
}
