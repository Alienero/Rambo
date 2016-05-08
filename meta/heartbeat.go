package meta

import (
	"time"

	"github.com/golang/glog"
)

// CreateAndHeartBeat create a k/v in etcd. And half a ttl, is will refresh the kv's ttl. If fail, will return
// an error. If error is nil it will return a named Stop function, it will stop the update tick.
func (m *Info) CreateAndHeartBeat(key, value string, ttl uint64, isThrow bool) (stop func(), err error) {
	timeout := time.Duration(ttl) * time.Second
	_, err = m.Create(key, value, ttl)
	if err != nil {
		return nil, err
	}
	inter := timeout / 2
	glog.Infof("CreateAndHeartBeat interval(%v)", inter)
	tick := time.NewTicker(inter)
	go func() {
		for {
			_, ok := <-tick.C
			if !ok {
				// stop
				glog.Info("tick is stoped")
				return
			}
			// update
			_, err := m.Update(key, value, ttl)
			if err != nil && isThrow {
				panic(err)
			}
		}
	}()
	return tick.Stop, nil
}
