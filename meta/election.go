package meta

import (
	"time"

	"sync"

	"github.com/coreos/go-etcd/etcd"
	"github.com/golang/glog"
)

// Election is elect leader in etcd
type Election struct {
	client     *etcd.Client
	localValue string
	dir        string
	ttl        uint64
	stop       chan bool
	isClosed   bool
	onece      sync.Once
	receiver   chan *etcd.Response
	// when is elect to master callback
	cb func(key string)
}

func NewElection(machines []string, dir string, ttl uint64, cb func(key string), localValue string) *Election {
	return &Election{
		client:     etcd.NewClient(machines),
		localValue: localValue,
		ttl:        ttl,
		receiver:   make(chan *etcd.Response, 1000),
		cb:         cb,
		dir:        dir,
	}
}

// GetMaster get master node
// if master node is not exist, it will elect new one
func (e *Election) GetMaster(key string) (string, error) {
renew:
	resp, err := e.client.Get(key, false, false)
	if err != nil {
		if etcdErr, ok := err.(*etcd.EtcdError); ok {
			if etcdErr.ErrorCode == NotFound {
				if _, err = e.client.Create(key, e.localValue, e.ttl); err != nil {
					goto renew
				}
				// this node is master
				e.cb(key)
				go e.update(key)
				return e.localValue, nil
			}
		}
		return "", err
	}
	return resp.Node.Value, nil
}

// Watch will watch keys with e.key perfix
// one instance only will called once
func (e *Election) Watch() {
	e.onece.Do(e.watch)
}

func (e *Election) watch() {
	_, err := e.client.Watch(e.dir, 0, true, e.receiver, e.stop)
	if err != nil {
		glog.Fatal(err)
	}
	for {
		resp, ok := <-e.receiver
		if !ok {
			glog.Info("masters watcher closed")
			return
		}
		if resp.Action == Expire {
			go e.electMaster(resp.Node.Key, resp.Node.Value)
		}
	}
}

func (e *Election) electMaster(key, lastValue string) {
renew:
	_, err := e.client.CompareAndSwap(key, e.localValue, e.ttl, lastValue, 0)
	if err == nil {
		e.cb(key)
		e.update(key)
	} else {
		// check master
		_, err := e.client.Get(key, false, false)
		if err != nil {
			time.Sleep(1 * time.Second)
			goto renew
		}
	}
}

// Stop will stop Election
func (e *Election) Stop() {
	// TOOD: impl
	e.isClosed = true
}

func (e *Election) update(key string) {
	t := time.NewTicker(time.Duration(e.ttl / 2))
	for {
		<-t.C
		// update node
		_, err := e.client.CompareAndSwap(key, e.localValue, e.ttl, e.localValue, 0)
		if err != nil {
			return
		}
	}
}
