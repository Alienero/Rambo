package meta

import "github.com/coreos/go-etcd/etcd"

// Election is elect leader in etcd
type Election struct {
	client     *etcd.Client
	key        string
	lastValue  string
	localValue string
	ttl        uint64
	stop       chan bool
	receiver   chan *etcd.Response
}

func (e *Election) initWatch() error {
	_, err := e.client.Watch(e.key, 0, false, e.receiver, e.stop)
	return err
}

func (e *Election) watch() error {
renew:
	// get the leader
	resp, err := e.client.Get(e.key, false, false)
	if err != nil {
		if etcdErr, ok := err.(*etcd.EtcdError); ok {
			// if leader not exist
			if etcdErr.ErrorCode == NotFoud {
				if err = e.elect(true); err == nil {
					goto renew
				}
			}
		}
		return err
	}
	e.lastValue = resp.Node.Value
	// watch leader node
	for {
		resp, ok := <-e.receiver
		if !ok {
			return nil
		}
		switch resp.Action {
		case Get, Update:
			// nothing to do
		case Delete:
			// stop
			e.stopElect()
			return nil
		case Expire:
			// elect new leader
			e.elect(false)
			goto renew
		case Set, CAS, Create:
			// new set
			goto renew
		}
	}
}

func (e *Election) elect(isCreate bool) (err error) {
	if isCreate {
		_, err = e.client.Create(e.key, e.localValue, e.ttl)
	} else {
		_, err = e.client.CompareAndSwap(e.key, e.localValue, e.ttl, e.lastValue, 0)
	}
	return
}

func (e *Election) stopElect() {
	// TODO: impl it
}
