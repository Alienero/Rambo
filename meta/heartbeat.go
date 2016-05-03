package meta

import "time"

// CreateAndHeartBeat create a k/v in etcd. And half a ttl, is will refresh the kv's ttl. If fail, will return
// an error. If error is nil it will return a named Stop function, it will stop the update tick.
func (m *Info) CreateAndHeartBeat(key, value string, ttl uint64, isThrow bool) (stop func(), err error) {
	_, err = m.Create(key, value, ttl)
	if err != nil {
		return nil, err
	}
	tick := time.NewTicker(time.Duration(ttl / 2))
	go func() {
		for {
			_, ok := <-tick.C
			if !ok {
				// stop
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
