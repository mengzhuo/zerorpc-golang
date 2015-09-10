package zerorpc

import (
	"net/rpc"
	"testing"
	"time"
)

type Args struct {
	X, Y int
}

type Calculator struct{ ch chan bool }

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	t.ch <- true
	return nil
}

func TestServerEndpoint(t *testing.T) {

	ch := make(chan bool)
	cal := &Calculator{ch}
	server := rpc.NewServer()
	server.Register(cal)

	codec := ServeEndpoint("inproc://zpc-server-endpoint")

	go server.ServeCodec(codec)
	ticker := time.NewTicker(500 * time.Microsecond)

	select {
	case <-cal.ch:
		codec.Close()
	case <-ticker.C:
		t.Errorf("Timeouted on ServeEndpoint")
	}
}

/*
func TestNewCodec(t *testing.T) {

	cal := new(Calculator)
	conn, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		t.Error(err)
	}
	server := rpc.NewServer()
	server.Register(cal)

	go server.ServeCodec(NewCodec(conn))
}
*/
