package kinesis

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awskinesis "github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

// KinesisStreamer -
type KinesisStreamer struct {
	stream *awskinesis.Kinesis
	name   string
}

// NewKinesisStreamer -
func NewKinesisStreamer(opts ...KinesisStreamerOption) (streamer.Streamer, error) {
	result := KinesisStreamer{}
	for _, opt := range opts {
		if err := opt(&result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Push implementes Streamer interface
func (streamer KinesisStreamer) Push(ctx context.Context, records []streamer.Record) error {

	kinesisEntries := []*awskinesis.PutRecordsRequestEntry{}

	for _, record := range records {
		kinesisEntries = append(kinesisEntries, &awskinesis.PutRecordsRequestEntry{
			Data:         record.Data,
			PartitionKey: aws.String(record.AggregateID),
		})
	}

	input := &awskinesis.PutRecordsInput{
		Records: kinesisEntries,
	}

	output, err := streamer.stream.PutRecordsWithContext(ctx, input)

	if err != nil {
		return err
	}
	if *output.FailedRecordCount != 0 {
		return errors.New("unable to push records")
	}

	return nil
}
