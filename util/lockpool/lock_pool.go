package lockpool

import (
	"sync"
	"sync/atomic"
)

// Pool is a lock pool
type Pool struct {
	mutex sync.RWMutex
	locks map[string]*Lock
}

// Lock is a lock pool's lock
type Lock struct {
	lock  *sync.RWMutex
	refer int64
}

// New get a lock pool instance
func New() *Pool {
	return &Pool{
		locks: make(map[string]*Lock),
	}
}

// Lock RWMutex lock
func (p *Pool) Lock(key string) {
	p.getLock(key).lock.Lock()
}

// Unlock RWMutex unlock
func (p *Pool) Unlock(key string) {
	p.getLock(key).lock.Unlock()
}

// RLock RWMutex RLock
func (p *Pool) RLock(key string) {
	l := p.getLock(key)
	l.lock.RLock()
	p.putLock(key, l)
}

// RUnlock RWMutex Runlock
func (p *Pool) RUnlock(key string) {
	l := p.getLock(key)
	l.lock.RUnlock()
	p.putLock(key, l)
}

func (p *Pool) getLock(key string) *Lock {
	p.mutex.RLock()
	l, ok := p.locks[key]
	if ok {
		atomic.AddInt64(&l.refer, 1)
		p.mutex.RUnlock()
		return l
	}
	p.mutex.Unlock()
	p.mutex.Lock()
	defer p.mutex.Unlock()
	l, ok = p.locks[key]
	if !ok {
		l = &Lock{
			lock:  new(sync.RWMutex),
			refer: 1,
		}
		p.locks[key] = l
	} else {
		atomic.AddInt64(&l.refer, 1)
	}
	return l
}

func (p *Pool) putLock(key string, l *Lock) {
	if atomic.AddInt64(&l.refer, -1) == 0 {
		p.mutex.Lock()
		if l.refer == 0 {
			delete(p.locks, key)
		}
		p.mutex.Unlock()
	}
}
