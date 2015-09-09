package zerorpc

import "fmt"

//go:generate msgp
//msgp:tuple ServerRequest
//msgp:tuple ServerResponse

type EventHeader struct {
	Id         string `msg:"message_id"`
	Version    int    `msg:"v"`
	ResponseTo string `msg:"response_to,omitempty"`
}

type ServerRequest struct {
	Header   *EventHeader
	Name     string
	Params   []interface{}
	Identity string `msg:"-"`
}

func (s *ServerRequest) String() string {
	return fmt.Sprintf("ID:%s Name:%s Args:%v", s.Header.Id, s.Name, s.Params)
}

func (s *ServerRequest) reset() {
	s.Header = nil
	s.Name = ""
	s.Params = nil
}

type ServerResponse struct {
	Header *EventHeader
	Name   string
	Params []interface{}
}
