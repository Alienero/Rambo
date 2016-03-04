package ddl

import (
	"sync"
	"sync/atomic"
	"time"
)

// Waitter is a Lock manage interface.
// key is [User]/[DataBase]/[Table]
type Waitter interface {
	Wait(key []string)
	Continue(key []string)

	Add(key []string)
	Done(key []string)
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
}

// NewWait get a default Wait pointer
func NewWait() *Wait {
	return &Wait{
		users:     make(Map),
		databases: make(Map),
		tables:    make(Map),
	}
}

// Wait lock a source use write lock.
func (w *Wait) Wait(key []string) {
	switch len(key) {
	case 1:
		w.getLock(key[0], &w.users).Lock()
	case 2:
		w.getLock(key[1], &w.databases).Lock()
	case 3:
		w.getLock(key[2], &w.tables).Lock()
	}
}

// Continue unlock a source use write lock.
func (w *Wait) Continue(key []string) {
	switch len(key) {
	case 1:
		w.getLock(key[0], &w.users).Lock()
	case 2:
		w.getLock(key[1], &w.databases).Lock()
	case 3:
		w.getLock(key[2], &w.tables).Lock()
	}
}

// Add a source lock.
// Should first lock big lock, then lock smaller.
func (w *Wait) Add(key []string) {
	if len(key) > 0 {
		w.getLock(key[0], &w.users).RLock()
	}
	if len(key) > 1 {
		w.getLock(key[1], &w.databases).RLock()
	}
	if len(key) > 2 {
		w.getLock(key[2], &w.tables).RLock()
	}
}

// Done release a lock source.
// Should first release small source then release bigger.
func (w *Wait) Done(key []string) {
	if len(key) > 2 {
		w.getLock(key[2], &w.tables).RUnlock()
	}
	if len(key) > 1 {
		w.getLock(key[1], &w.databases).RUnlock()
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
