package zerorpc

import (
	"net/rpc"
	"testing"
	"time"

	"github.com/golang/glog"
)

type Args struct {
	X, Y int
}

type Calculator struct{}

func (t *Calculator) Add(args *Args, reply *int) error {
	glog.Error(args, reply)
	*reply = args.X + args.Y
	//reply = &Args{args.Y, args.X}
	return nil
}

func TestServer(t *testing.T) {

	cal := new(Calculator)

	server := rpc.NewServer()
	server.Register(cal)

	codec := ServeEndpoint("tcp://*:9999")

	go server.ServeCodec(codec)

	time.Sleep(2 * time.Second)
	codec.Close()
}
