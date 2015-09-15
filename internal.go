package zerorpc

import "github.com/golang/glog"

type ZArgs []int

func NewInternalService() *InternalService {
	i := &InternalService{make([][]string, 0)}
	i.registerDoc("zerorpc", "Demo")
	return i
}

type InternalService struct {
	methods [][]string
}

func (z *InternalService) registerDoc(name, doc string) {
	z.methods = append(z.methods, []string{name, "", doc})
}

func (z *InternalService) Inspect(args *ZArgs, reply *map[string]interface{}) error {
	*reply = map[string]interface{}{"methods": z.methods}
	return nil
}
func (z *InternalService) HeartBeat(args *ZArgs, reply *int) error {
	glog.Error(args, reply)
	return nil
}
