package streamer

type Record struct {
	AggregateID string `json:"id"`
	Type        string `json:"t"`
	Data        []byte `json:"d"`
}
