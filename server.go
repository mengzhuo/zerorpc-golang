package zerorpc

import (
	"fmt"
	"net/rpc"
	"reflect"
	"sync"
	"time"

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
	channel map[string]*channel
}

func (c *serverCodec) Close() error {
	return c.zsock.Close()
}

func (c *serverCodec) checkIdentity(identity string) {

	if ch, ok := c.channel[identity]; !ok {
		ch = &channel{c, identity, time.NewTicker(5 * time.Second), 1, make(chan bool)}
		c.channel[identity] = ch
		go ch.run()
	} else {
		ch.counter += 1
	}
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) (err error) {

	c.req.reset()

	msg, err := c.zsock.RecvMessageBytes(0)
	if err != nil || len(msg) < 2 {
		return err
	}
	identity := string(msg[0])
	//XXX only last one?
	o, err := c.req.UnmarshalMsg(msg[len(msg)-1])

	if err != nil || c.req.Name == "" {
		glog.Errorf("o=%#v err=%s", o, err)
		return fmt.Errorf("zerorpc: Error %s c.req=%#v leftbytes=%#v", err.Error(), c.req, o)
	}

	if c.req.Header.Version != PROTOCAL_VERSION {
		return fmt.Errorf("zerorpc: Version not matching with request, expecting %d but sending %d", PROTOCAL_VERSION, c.req.Header.Version)
	}

	switch c.req.Name {
	case "_zpc_hb":
		c.req.Name = "InternalService.HeartBeat"
	case "_zerorpc_inspect":
		c.req.Name = "InternalService.Inspect"
	}

	glog.V(1).Infof("Receiving Event %s", c.req)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.checkIdentity(identity)
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

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := val.Field(i)
			if f.CanSet() {
				f.Set(reflect.ValueOf(c.req.Params[i]).Convert(f.Type()))
			} else {
				return fmt.Errorf("Field :%s can't be set", f)
			}
		}
	default:
		x = val.Interface()
	}

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

	if ch, ok := c.channel[b.Identity]; ok {
		ch.counter -= 1
		if ch.counter <= 0 {
			select {
			case ch.closed <- true:
			default:
			}
			delete(c.channel, b.Identity)
		}
	}
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
		params[0] = body
		name = "OK"
	}

	resp := &ServerResponse{Header: b.Header, Name: name, Params: params}

	o, err := resp.MarshalMsg(nil)
	if err != nil {
		glog.Error(err, o)
	}

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
		channel: make(map[string]*channel),
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
