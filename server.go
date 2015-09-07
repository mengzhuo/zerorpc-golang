package zerorpc

import (
	"io"
	"net/rpc"
	"sync"

	"github.com/golang/glog"
	"github.com/ugorji/go/codec"
	"github.com/zeromq/goczmq"
)

var mh codec.MsgpackHandle

// ZeroRPC protocol version
const PROTOCAL_VERSION = 3

type RequestHeader struct {
	Id         string `codec:"message_id"`
	Version    int    `codec:"v"`
	ResponseTo string `codec:"response_to"`
}

type serverRequest struct {
	Header *RequestHeader
	Name   string                        `codec:"name,omitempty"`
	Args   codec.MsgpackSpecRpcMultiArgs `codec:"args,omitempty"`
}

func (r *serverRequest) reset() {

	r.Name = ""
}

type serverResponse struct {
	Header *RequestHeader
	Name   string `codec:"name,omitempty"`
	Args   codec.MsgpackSpecRpcMultiArgs
}

type serverCodec struct {
	dec  *codec.Decoder // for reading msgpack values
	enc  *codec.Encoder // for writing msgpack values
	conn io.ReadWriter
	seq  uint64
	req  *serverRequest

	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]string
}

func (c *serverCodec) Close() error {

	return nil
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {

	var vv interface{}
	if err := c.dec.Decode(&vv); err != nil {
		glog.Info(vv)
		glog.Error(err)
		return err
	}
	glog.Info(vv)
	c.mutex.Lock()
	defer c.mutex.Unlock()

	r.ServiceMethod = c.req.Name
	c.seq++
	c.pending[c.seq] = c.req.Header.Id
	r.Seq = c.seq
	glog.Info(r)
	glog.Info(c)
	return nil
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {

	glog.Info("RRB", x)
	if x == nil {
		return nil
	}

	return nil
}

func (c *serverCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {

	return nil
}

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
func NewServerCodec(conn io.ReadWriter) rpc.ServerCodec {

	r := &serverRequest{&RequestHeader{}, "", make([]interface{}, 0)}
	return &serverCodec{
		dec:     codec.NewDecoder(conn, &mh),
		enc:     codec.NewEncoder(conn, &mh),
		conn:    conn,
		pending: make(map[uint64]string),
		req:     r,
	}
}

func NewServer(address string) rpc.ServerCodec {

	raw_sock, err := goczmq.NewRouter(address)
	if err != nil {
		panic(err)
	}
	return NewServerCodec(raw_sock)
}
