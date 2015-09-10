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

	params := make([]interface{}, 1, 1)

	//ele := reflect.ValueOf(body).Elem()
	rb := reflect.New(reflect.TypeOf(body).Elem())
	glog.Errorf("b=%#v kind=%s be=%#v kind=%s", reflect.TypeOf(body), reflect.TypeOf(body).Kind(),
		reflect.TypeOf(body).Elem(), reflect.TypeOf(body).Elem().Kind())

	glog.Errorf("rb=%#v kind=%s rbe=%#v kind=%s", reflect.ValueOf(body), reflect.ValueOf(body).Kind(),
		reflect.ValueOf(body).Elem(), reflect.ValueOf(body).Elem().Kind())

	rb.Set(reflect.ValueOf(body).Addr())

	params[0] = rb

	/*
		switch ele.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64:
			params[0] = ele.Interface().(int)
		case reflect.String:
			params[0] = ele.Interface().(string)
		default:
			glog.Errorf("Kind =%s valid=%b nil=%b", ele.Kind(), ele.IsValid(), ele.IsNil())
		}
		glog.Errorf("%#v, %#v, %#v", r, ele)
	*/
	resp := &ServerResponse{Header: b.Header, Name: "OK", Params: params}

	if r.Error != "" {
		resp.Name = "ERR"
		resp.Params = []interface{}{r.Error}
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
