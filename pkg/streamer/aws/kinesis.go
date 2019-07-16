package aws

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

// KinesisStreamer -
type KinesisStreamer struct {
	stream *kinesis.Kinesis
	name   string
}

// NewKinesisStreamer -
func NewKinesisStreamer() streamer.Streamer {

	sess := session.Must(session.NewSession())

	stream := kinesis.New(sess, &aws.Config{
		Region: aws.String(endpoints.EuWest1RegionID),
	})

	return KinesisStreamer{
		stream: stream,
	}
}

// Push implementes Streamer interface
func (streamer KinesisStreamer) Push(ctx context.Context, records []streamer.Record) error {

	kinesisEntries := []*kinesis.PutRecordsRequestEntry{}

	for _, record := range records {
		kinesisEntries = append(kinesisEntries, &kinesis.PutRecordsRequestEntry{
			Data:         record.Data,
			PartitionKey: aws.String(record.AggregateID),
		})
	}

	input := &kinesis.PutRecordsInput{
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
