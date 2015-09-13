package zerorpc

import (
	"net/rpc"
	"os/exec"
	"testing"
	"time"
)

type Args struct {
	X, Y int
}

type Calculator struct{ ch chan bool }

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func run(cmd ...string) chan bool {
	c := exec.Command("zerorpc", cmd...)
	result := make(chan bool)
	go func(c *exec.Cmd) {
		c.Run()
		result <- true
	}(c)
	return result
}

func TestServerEndpoint(t *testing.T) {

	ch := make(chan bool)
	cal := &Calculator{ch}
	server := rpc.NewServer()
	server.Register(cal)

	codec := ServeEndpoint("tcp://*:12345")
	cha1 := run("-j", "tcp://localhost:12345", "Calculator.Add", "1", "2")
	cha2 := run("-j", "tcp://localhost:12345", "Calculator.Add", "3", "4")
	go server.ServeCodec(codec)

	time.Sleep(500 * time.Millisecond)

	ticker := time.NewTicker(500 * time.Millisecond)
	go func(t *testing.T) {
		<-ticker.C
		t.Errorf("Timeouted on ServeEndpoint")
	}(t)
	<-cha1
	<-cha2
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
