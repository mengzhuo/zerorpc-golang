package zerorpc

import (
	"net/rpc"
	"sync"

	"github.com/golang/glog"
	zmq "github.com/pebbe/zmq4"
)

// ZeroRPC protocol version
const PROTOCAL_VERSION = 3

type serverCodec struct {
	zsock *zmq.Socket
	seq   uint64
	// temporary work space
	req ServerRequest

	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]string
}

func (c *serverCodec) Close() error {
	return c.zsock.Close()
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.req.reset()

	msg, err := c.zsock.RecvMessageBytes(0)
	if err != nil {
		return err
	}
	glog.Infof("MSG Len:%d -> %#v", len(msg), msg)

	o, err := c.req.UnmarshalMsg(msg[len(msg)-1])
	if err != nil {
		glog.Errorf("o=%#v err=%s", o, err)
		return err
	} else {
		glog.Infof("c.req=%#v o=%#v", c.req, o)
	}

	r.ServiceMethod = c.req.Name
	c.seq++
	c.pending[c.seq] = c.req.Header.Id
	r.Seq = c.seq

	return
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {
	return nil
}

func (c *serverCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {

	return nil
}

func NewSocket(address string) (zsock *zmq.Socket, err error) {

	zsock, err = zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		return
	}
	err = zsock.Bind(address)
	return
}

func ServeEndpoint(address string) rpc.ServerCodec {

	sock, err := NewSocket(address)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	return &serverCodec{
		zsock:   sock,
		seq:     0,
		pending: make(map[uint64]string),
		req:     ServerRequest{},
	}
}
