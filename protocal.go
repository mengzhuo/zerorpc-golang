package zerorpc

//go:generate msgp

type EventHeader struct {
	Id         string `msg:"message_id"`
	Version    int    `msg:"v"`
	ResponseTo string `msg:"response_to,omitempty"`
}

type ServerRequest struct {
	Header *EventHeader
	Name   string        `msg:"name"`
	Args   []interface{} `msg:"args"`
}

func (s *ServerRequest) reset() {
	s.Header = nil
	s.Name = ""
	s.Args = nil
}

type ServerResponse struct {
	Header *EventHeader
	Name   string        `msg:"name"`
	Args   []interface{} `msg:"args"`
}
