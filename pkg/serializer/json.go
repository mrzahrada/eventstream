package serializer

import (
	"encoding/json"
	"reflect"

	"github.com/mrzahrada/eventstream/pkg/event"
	"github.com/mrzahrada/eventstream/pkg/streamer"
)

type jsonEvent struct {
	Type string          `json:"t"`
	Data json.RawMessage `json:"d"`
}

// JSONSerializer -
type JSONSerializer struct {
	eventTypes map[string]reflect.Type
}

// NewJSONSerializer -
func NewJSONSerializer(events ...event.Event) Serializer {
	serializer := JSONSerializer{
		eventTypes: map[string]reflect.Type{},
	}
	serializer.Bind(events...)
	return serializer
}

// Bind registers the specified events with the serializer; may be called more than once
func (serializer *JSONSerializer) Bind(events ...event.Event) {
	for _, event := range events {
		eventType, t := eventType(event)
		serializer.eventTypes[eventType] = t
	}
}

// Marshal converts an event into its persistent type, Record
func (serializer JSONSerializer) Marshal(v event.Event) (streamer.Record, error) {

	eventType, _ := eventType(v)

	data, err := json.Marshal(v)
	if err != nil {
		// TODO: enrich ErrMarshalEvent
		return streamer.Record{}, err
	}

	data, err = json.Marshal(jsonEvent{
		Type: eventType,
		Data: json.RawMessage(data),
	})
	if err != nil {
		return streamer.Record{}, ErrMarshalEvent
	}

	return streamer.Record{
		AggregateID: v.AggregateID(),
		Data:        data,
		Type:        eventType,
	}, nil
}

// Unmarshal converts the persistent type, Record, into an Event instance
func (serializer JSONSerializer) Unmarshal(record streamer.Record) (event.Event, error) {

	wrapper := jsonEvent{}
	err := json.Unmarshal(record.Data, &wrapper)
	if err != nil {
		// TODO: enrich ErrUnmarshalEvent
		return nil, err
	}

	t, ok := serializer.eventTypes[wrapper.Type]
	if !ok {
		return nil, ErrUnknownEventType
	}

	v := reflect.New(t).Interface()
	err = json.Unmarshal(wrapper.Data, v)
	if err != nil {
		return nil, ErrUnmarshalEvent
	}

	return v.(event.Event), nil
}

func eventType(event event.Event) (string, reflect.Type) {
	t := reflect.TypeOf(event)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name(), t
}
