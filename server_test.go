package zerorpc

import (
	"net/rpc"
	"testing"
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

	codec := ServeEndpoint("tcp://*:9999")

	server.ServeRequest(codec)
}
