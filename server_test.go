package zerorpc

import (
	"net/rpc"
	"testing"

	"github.com/zeromq/goczmq"
)

type Args struct {
	X, Y int
}

type Calculator struct{}

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func TestServer(t *testing.T) {

	cal := new(Calculator)

	server := rpc.NewServer()
	server.Register(cal)

	router, e := goczmq.NewRouter("tcp://*:9999")
	if e != nil {
		t.Error(e)
	}

	server.ServeCodec(NewServerCodec(router))

}
