package eventbridge

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	eb "github.com/aws/aws-sdk-go/service/eventbridge"
)

type EventBridgeStreamerOption func(*EventBridgeStreamer) error

func WithEventBusName(name string) EventBridgeStreamerOption {
	return func(streamer *EventBridgeStreamer) error {
		streamer.ebName = name
		return nil
	}
}

func WithEventBridgeSource(source string) EventBridgeStreamerOption {
	return func(streamer *EventBridgeStreamer) error {
		streamer.source = source
		return nil
	}
}

func WithEventBridgeRegion(region string) EventBridgeStreamerOption {
	return func(streamer *EventBridgeStreamer) error {
		sess, err := session.NewSession()
		if err != nil {
			return err
		}
		stream := eb.New(sess, &aws.Config{
			// Region: aws.String(endpoints.EuWest1RegionID),
			Region: aws.String(region),
		})
		streamer.stream = stream
		return nil
	}
}
