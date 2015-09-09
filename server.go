package zerorpc

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/golang/glog"
	zmq "github.com/pebbe/zmq4"
	"github.com/satori/go.uuid"
)

// ZeroRPC protocol version
const PROTOCAL_VERSION = 3

type serverCodec struct {
	zsock *zmq.Socket
	seq   uint64
	// temporary work space
	req ServerRequest

	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]ServerRequest
}

func (c *serverCodec) Close() error {
	return c.zsock.Close()
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) (err error) {

	c.req.reset()

	msg, err := c.zsock.RecvMessageBytes(0)
	if err != nil {
		return err
	}
	identity := string(msg[0])

	//XXX only last one?
	o, err := c.req.UnmarshalMsg(msg[len(msg)-1])

	glog.Errorf("msg %#v", msg)

	if err != nil || c.req.Name == "" {
		glog.Errorf("o=%#v err=%s", o, err)
		return fmt.Errorf("zerorpc: Error %s c.req=%#v leftbytes=%#v", err.Error(), c.req, o)
	}

	if c.req.Header.Version != PROTOCAL_VERSION {
		return fmt.Errorf("zerorpc: Version not matching with request, expecting %d but sending %d", PROTOCAL_VERSION, c.req.Header.Version)
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	glog.Infof("zerorpc: req:%s", c.req)

	c.req.Identity = identity
	r.ServiceMethod = c.req.Name
	c.seq++
	c.pending[c.seq] = c.req
	r.Seq = c.seq

	return
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {
	// We already decoded params
	var params [1]interface{}
	params[0] = x
	return nil
}

func (c *serverCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {

	c.mutex.Lock()
	b, ok := c.pending[r.Seq]

	if !ok {
		c.mutex.Unlock()
		return fmt.Errorf("invalid sequence number:%d in response", r.Seq)
	}
	delete(c.pending, r.Seq)
	c.mutex.Unlock()

	b.Header.ResponseTo = b.Header.Id
	b.Header.Id = uuid.NewV4().String()

	glog.Error(r, body)

	resp := &ServerResponse{Header: b.Header, Name: "OK", Params: []interface{}{nil}}

	if r.Error != "" {
		resp.Name = "ERR"
		resp.Params = append(resp.Params, r.Error)
	}

	o, err := resp.MarshalMsg(nil)
	if err != nil {
		glog.Error(err, o)
	}

	glog.Errorf("zerorpc: resp:%s", resp)
	_, err = c.zsock.SendMessage(b.Identity, "", o)
	return err
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
		pending: make(map[uint64]ServerRequest),
		req:     ServerRequest{},
	}
}
