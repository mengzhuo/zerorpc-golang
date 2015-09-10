package zerorpc

import (
	"fmt"
	"net/rpc"
	"reflect"
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
	val := reflect.ValueOf(x)

	if !val.IsValid() {
		return fmt.Errorf("zerorpc: request is not valid! x=%#v", x)
	}

	t := reflect.TypeOf(x)

	if t.Kind() == reflect.Ptr {
		val = val.Elem()
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := val.Field(i)
		if f.CanSet() {
			f.Set(reflect.ValueOf(c.req.Params[i]).Convert(f.Type()))
		} else {
			return fmt.Errorf("Field :%s can't be set", f)
		}
	}

	return nil
}

func (c *serverCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {

	// body must be a pointer of Interface

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
	params := make([]interface{}, 1)
	name := "ERR"

	if r.Error != "" {
		glog.Error(r.Error)
		params[0] = r.ServiceMethod
		params = append(params, r.Error)
		params = append(params, r.Error)
	} else {

		ele := reflect.ValueOf(body)
		if ele.Kind() == reflect.Ptr {
			ele.Elem()
		}
		params[0] = ele.Interface()
		name = "OK"
	}

	resp := &ServerResponse{Header: b.Header, Name: name, Params: params}
	o, err := resp.MarshalMsg(nil)
	if err != nil {
		glog.Error(err, o)
	}
	glog.Errorf("zerorpc: resp:%s", resp)
	_, err = c.zsock.SendMessage(b.Identity, "", o)

	return err
}

func NewConn(address string) (zsock *zmq.Socket, err error) {

	zsock, err = zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		return
	}
	err = zsock.Bind(address)
	return
}

func NewCodec(conn *zmq.Socket) rpc.ServerCodec {

	return &serverCodec{
		zsock:   conn,
		seq:     0,
		pending: make(map[uint64]ServerRequest),
		req:     ServerRequest{},
	}

}

func ServeEndpoint(address string) rpc.ServerCodec {

	sock, err := NewConn(address)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	return NewCodec(sock)
}
