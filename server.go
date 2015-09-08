package zerorpc

import (
	"net/rpc"
	"sync"

	"github.com/golang/glog"
	zmq "github.com/pebbe/zmq4"
	"github.com/ugorji/go/codec"
)

var mh codec.MsgpackHandle

// ZeroRPC protocol version
const PROTOCAL_VERSION = 3

type EventHeader struct {
	Id         string `codec:"message_id"`
	Version    int    `codec:"v"`
	ResponseTo string `codec:"response_to,omitempty"`
}

type serverRequest struct {
	Header *EventHeader
	Name   string
	Args   codec.MsgpackSpecRpcMultiArgs
}

func (s *serverRequest) reset() {
	s.Header = nil
	s.Name = ""
	s.Args = nil
}

type serverResponse struct {
	Header *EventHeader
	Name   string
	Args   codec.MsgpackSpecRpcMultiArgs
}

type serverCodec struct {
	zsock *zmq.Socket
	seq   uint64
	// temporary work space
	req serverRequest

	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]string
	buf     []byte
	dec     *codec.Decoder
}

func (c *serverCodec) Close() error {
	return c.zsock.Close()
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.buf = c.buf[:0]
	defer func(c *serverCodec) {
		c.buf = c.buf[:0]
	}(c)

	msg, err := c.zsock.RecvMessageBytes(0)
	glog.Infof("MSG Len:%d -> %s", len(msg), msg)

	dec := codec.NewDecoderBytes(msg[len(msg)-1], &mh)

	if err = dec.Decode(&c.req); err != nil {
		glog.Errorf("Decode Error:%s bytes=%s", err, c.buf)
		return
	}
	glog.Infof("%#v HEADER:%#v", c.req, c.req.Header)
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
	buf := []byte{}
	dec := codec.NewDecoderBytes(buf, &mh)
	return &serverCodec{
		zsock:   sock,
		seq:     0,
		buf:     buf,
		dec:     dec,
		pending: make(map[uint64]string),
	}
}
