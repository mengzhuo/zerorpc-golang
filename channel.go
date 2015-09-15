package zerorpc

import (
	"time"

	"github.com/golang/glog"
	"github.com/satori/go.uuid"
)

type channel struct {
	codec    *serverCodec
	identity string
	ticker   *time.Ticker
	counter  int // count pending req
	closed   chan bool
}

func (c *channel) run() {

	head := &EventHeader{Id: uuid.NewV4().String(), Version: PROTOCAL_VERSION}
	req := &ServerResponse{head, "_zpc_hb", make([]interface{}, 0)}
	o, err := req.MarshalMsg(nil)
	if err != nil {
		glog.Error(err)
		return
	}

	for {
		select {
		case <-c.ticker.C:
			c.codec.mutex.Lock()
			glog.Error(o)
			_, err = c.codec.zsock.SendMessageDontwait(c.identity, o)
			if err != nil {
				c.ticker.Stop()
				c.codec = nil
				c.counter = 0
				glog.Error(err)
			}
			c.codec.mutex.Unlock()
			return
		case <-c.closed:
			c.ticker.Stop()
			c.codec = nil
			c.ticker = nil
			return
		}
	}
}
