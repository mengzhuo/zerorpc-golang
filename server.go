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

type EventHeader struct {
	Id      string `codec:"message_id"`
	Version int    `codec:"v"`
}

type serverRequest struct {
	*EventHeader
	Name string
	Args codec.MsgpackSpecRpcMultiArgs
}

func (r *serverRequest) reset() {

	r.Id = ""
	r.Version = 0
	r.Name = ""
	r.Args = nil

}

type serverResponse struct {
	*EventHeader
	ResponseTo string `codec:"response_to"`
	Args       codec.MsgpackSpecRpcMultiArgs
}

type serverCodec struct {
	dec  *codec.Decoder // for reading msgpack values
	enc  *codec.Encoder // for writing msgpack values
	conn io.ReadWriter
	seq  uint64
	// temporary work space
	req *serverRequest

	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]string
}

func (c *serverCodec) Close() error {

	return nil
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {

	glog.Info(r)
	c.req.reset()

	if err := c.dec.Decode(c.req); err != nil {
		glog.Error(err)
		return err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()

	r.ServiceMethod = c.req.Name
	c.seq++
	c.pending[c.seq] = c.req.Id
	r.Seq = c.seq
	glog.Info(r)
	glog.Info(c)
	return nil
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {

	glog.Info(x)
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
	return &serverCodec{
		dec:     codec.NewDecoder(conn, &mh),
		enc:     codec.NewEncoder(conn, &mh),
		conn:    conn,
		pending: make(map[uint64]string),
		req:     &serverRequest{},
	}
}

func NewServer(address string) rpc.ServerCodec {

	conn, err := goczmq.NewRouter(address)
	if err != nil {
		panic(err)
	}
	return NewServerCodec(conn)
}
