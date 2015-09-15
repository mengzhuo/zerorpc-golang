package zerorpc

import "time"

type channel struct {
	ticker *time.Ticker
	closed chan bool
}
