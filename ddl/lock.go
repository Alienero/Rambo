package ddl

import (
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Alienero/Rambo/util/lockpool"
)

// TODO: add locks DDLLock gc

// Waitter is a Lock manage interface.
// key is [User]/[DataBase]/[Table]
type Waitter interface {
	Wait(dl *DDLLock)
	Continue(dl *DDLLock)

	Add(key string)
	Done(key string)
}

// Map mapping key and lock
type Map map[string]*Lock

// Lock include a lock and refer count
type Lock struct {
	lock  sync.RWMutex
	refer int64
}

// Lock lock a write lock
func (l *Lock) Lock() {
	atomic.AddInt64(&l.refer, 1)
	l.lock.Lock()
}

// UnLock unlock a write lock
func (l *Lock) UnLock() int64 {
	ret := atomic.AddInt64(&l.refer, -1)
	l.lock.Unlock()
	return ret
}

// RLock lock a read lock
func (l *Lock) RLock() {
	atomic.AddInt64(&l.refer, 1)
	l.lock.RLock()
}

// RUnlock unlock a read lock
func (l *Lock) RUnlock() int64 {
	ret := atomic.AddInt64(&l.refer, -1)
	l.lock.RUnlock()
	return ret
}

// Wait is default Waitter impl
type Wait struct {
	users     Map
	databases Map
	tables    Map
	sync.RWMutex

	// locks key
	locks map[string]*DDLLock
	lpool *lockpool.Pool
}

// NewWait get a default Wait pointer
func NewWait() *Wait {
	return &Wait{
		users:     make(Map),
		databases: make(Map),
		tables:    make(Map),
		locks:     make(map[string]*DDLLock),
		lpool:     lockpool.New(),
	}
}

// Wait lock a source use write lock.
// for rpc
func (w *Wait) Wait(dl *DDLLock) {
	dl.status = StatusLock
	w.lpool.RLock(dl.ID)
	l, ok := w.locks[dl.ID]
	if ok && w.isPass(dl, l) {
		w.lpool.RUnlock(dl.ID)
		return
	}
	w.lpool.RUnlock(dl.ID)
	w.lpool.Lock(dl.ID)
	defer w.lpool.Unlock(dl.ID)
	l, ok = w.locks[dl.ID]
	if ok && w.isPass(dl, l) {
		return
	}
	// create a new Lock
	key := strings.Split(dl.Key, "/")
	switch len(key) {
	case 1:
		w.getLock(key[0], &w.users).Lock()
	case 2:
		w.getLock(path.Join(key[0], key[1]), &w.databases).Lock()
	case 3:
		w.getLock(path.Join(key[0], key[1], key[2]), &w.tables).Lock()
	}
	w.locks[dl.ID] = dl
}

// Continue unlock a source use write lock.
// for rpc
func (w *Wait) Continue(dl *DDLLock) {
	dl.status = StatusUnlock
	w.lpool.RLock(dl.ID)
	l, ok := w.locks[dl.Key]
	if !ok {
		w.lpool.RUnlock(dl.ID)
		return
	} else if w.isPass(dl, l) {
		w.lpool.RUnlock(dl.ID)
		return
	}
	w.lpool.RUnlock(dl.ID)

	// unlock
	key := strings.Split(dl.Key, "/")
	switch len(key) {
	case 1:
		w.getLock(key[0], &w.users).Lock()
	case 2:
		w.getLock(path.Join(key[0], key[1]), &w.databases).Lock()
	case 3:
		w.getLock(path.Join(key[0], key[1], key[2]), &w.tables).Lock()
	}

	w.lpool.Lock(dl.ID)
	l.status = StatusUnlock
	w.lpool.Unlock(dl.ID)
}

// true is pass
func (w *Wait) isPass(remote *DDLLock, local *DDLLock) bool {
	if remote.Version < local.Version {
		return true
	}
	if remote.ID == local.ID && remote.status == remote.status {
		return true
	}
	return false
}

// Add a source lock.
// Should first lock big lock, then lock smaller.
func (w *Wait) Add(keys string) {
	key := strings.Split(keys, "/")
	if len(key) > 0 {
		w.getLock(key[0], &w.users).RLock()
	}
	if len(key) > 1 {
		w.getLock(path.Join(key[0], key[1]), &w.databases).RLock()
	}
	if len(key) > 2 {
		w.getLock(path.Join(key[0], key[1], key[2]), &w.tables).RLock()
	}
}

// Done release a lock source.
// Should first release small source then release bigger.
func (w *Wait) Done(keys string) {
	key := strings.Split(keys, "/")
	if len(key) > 2 {
		w.getLock(path.Join(key[0], key[1], key[2]), &w.tables).RUnlock()
	}
	if len(key) > 1 {
		w.getLock(path.Join(key[0], key[1]), &w.databases).RUnlock()
	}
	if len(key) > 0 {
		w.getLock(key[0], &w.users).RUnlock()
	}
}

// GCLock gc unuseful lock
func (w *Wait) GCLock(inter time.Duration) {
	t := time.NewTicker(inter)
	go func() {
		defer t.Stop()
		for {
			select {
			case <-t.C:
				w.gcCheck(&w.users)
				w.gcCheck(&w.databases)
				w.gcCheck(&w.tables)
			}
		}
	}()
}

func (w *Wait) gcCheck(m *Map) {
	for k, v := range *m {
		if atomic.LoadInt64(&v.refer) == 0 {
			w.Lock()
			delete(*m, k)
			w.Unlock()
		}
	}
}

func (w *Wait) getLock(key string, m *Map) *Lock {
	w.RLock()
	defer w.RUnlock()
	lock := (*m)[key]
	if lock == nil {
		w.RUnlock()
		w.Lock()
		if lock = (*m)[key]; lock == nil {
			lock = new(Lock)
			(*m)[key] = lock
		}
		w.Unlock()
		w.RLock()
	}
	return lock
}

var _ Waitter = NewWait()
