package kinesis

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awskinesis "github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisStreamerOption func(*KinesisStreamer) error

func WithKinesisStreamName(name string) KinesisStreamerOption {
	return func(streamer *KinesisStreamer) error {
		streamer.name = name
		return nil
	}
}

func WithKinesisReagion(region string) KinesisStreamerOption {
	return func(streamer *KinesisStreamer) error {
		sess, err := session.NewSession()
		if err != nil {
			return err
		}

		streamer.stream = awskinesis.New(sess, &aws.Config{
			Region: aws.String(region),
		})
		return nil
	}
}
