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
	server.RegisterName("_zerorpc_inspect", cal)

	router, e := goczmq.NewRouter("tcp://*:9999")
	if e != nil {
		t.Error(e)
	}

	codec := NewServerCodec(router)

	for {
		server.ServeRequest(codec)
	}
}
