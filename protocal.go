package zerorpc

//go:generate msgp
//msgp:tuple ServerRequest
//msgp:tuple ServerResponse

type EventHeader struct {
	Id         string `msg:"message_id"`
	Version    int    `msg:"v"`
	ResponseTo string `msg:"response_to,omitempty"`
}

type ServerRequest struct {
	Header *EventHeader
	Name   string
	Args   []interface{} `msg:"args,omitempty"`
}

func (s *ServerRequest) reset() {
	s.Header = nil
	s.Name = ""
	s.Args = nil
}

type ServerResponse struct {
	Header *EventHeader
	Name   string
	Args   []interface{} `msg:"args,omitempty"`
}
