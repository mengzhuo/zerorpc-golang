package zerorpc

import (
	"bytes"
	"fmt"
	"net/rpc"
	"sync"

	"github.com/golang/glog"
	zmq "github.com/pebbe/zmq4"
	"gopkg.in/vmihailenco/msgpack.v2"
)

// ZeroRPC protocol version
const PROTOCAL_VERSION = 3

type serverCodec struct {
	zsock *zmq.Socket
	seq   uint64
	// temporary work space
	req []interface{}

	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]string
	buf     *bytes.Buffer
	dec     *msgpack.Decoder
}

func (c *serverCodec) Close() error {
	return c.zsock.Close()
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.buf.Reset()
	c.req = c.req[:0]

	msg, err := c.zsock.RecvMessageBytes(0)
	if err != nil {
		return err
	}
	glog.Infof("MSG Len:%d -> %#v", len(msg), msg)
	c.buf.Write(msg[len(msg)-1])

	if err = c.dec.Decode(&c.req); err != nil {
		glog.Errorf("Decode Error:%s bytes=%s", err, msg)
		return
	}
	glog.Infof("%#v", c.req)

	r.ServiceMethod = c.req[1].(string)
	c.seq++
	c.pending[c.seq] = fmt.Sprintf("%#v", c.req[0].(map[interface{}]interface{}))
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
	buf := bytes.NewBuffer([]byte{})
	return &serverCodec{
		zsock:   sock,
		seq:     0,
		pending: make(map[uint64]string),
		buf:     buf,
		dec:     msgpack.NewDecoder(buf),
	}
}
