package zerorpc

import "time"

type channel struct {
	ticker  *time.Ticker
	counter int // count pending req
	closed  chan bool
}
