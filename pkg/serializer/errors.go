package serializer

import "errors"

var (
	// ErrMarshalEvent -
	ErrMarshalEvent = errors.New("unable to marshal event")

	// ErrUnmarshalEvent -
	ErrUnmarshalEvent = errors.New("unable to unmarshal event")

	// ErrUnknownEventType -
	ErrUnknownEventType = errors.New("unknown event type")
)
