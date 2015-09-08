package zerorpc

// go:generate msgp

type EventHeader struct {
	Id         string `msgpack:"message_id"`
	Version    int    `msgpack:"v"`
	ResponseTo string `msgpack:"response_to,omitempty"`
}

type serverRequest struct {
	Header *EventHeader
	Name   string
	Args   []interface{}
}

func (s *serverRequest) reset() {
	s.Header = nil
	s.Name = ""
	s.Args = nil
}

type serverResponse struct {
	Header *EventHeader
	Name   string
	Args   []interface{}
}
