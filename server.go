package zerorpc

import (
	"io"
	"net/rpc"
	"sync"

	zmq "github.com/pebbe/zmq4"
	"github.com/ugorji/go/codec"
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

type serverResponse struct {
	*EventHeader
	ResponseTo string `codec:"response_to"`
	Args       codec.MsgpackSpecRpcMultiArgs
}

type serverCodec struct {
	dec *codec.Decoder // for reading msgpack values
	enc *codec.Encoder // for writing msgpack values
	c   io.Closer

	// temporary work space
	req serverRequest

	mutex   sync.Mutex // protects seq, pending
	pending map[string]*EventHeader
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {

	return nil
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {

	return nil
}

func (c *serverCodec) Close() error {

	return nil
}

func (c *serverCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {

	return nil
}

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return &serverCodec{
		dec:     codec.NewDecoder(conn, &mh),
		enc:     codec.NewEncoder(conn, &mh),
		c:       conn,
		pending: make(map[string]*EventHeader),
	}
}

type socket struct {
	*zmq.Socket
}
