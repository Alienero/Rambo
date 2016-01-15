package safemap

import (
	"sync"
)

type Map struct {
	lock    *sync.RWMutex
	element map[interface{}]interface{}
}

func NewMap() *Map {
	return &Map{new(sync.RWMutex), make(map[interface{}]interface{})}
}

func (s *Map) Get(key interface{}) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.element[key]
}

func (s *Map) Set(key interface{}, vaule interface{}) {
	s.lock.Lock()
	s.element[key] = vaule
	s.lock.Unlock()
}

func (s *Map) Delete(key interface{}) {
	s.lock.Lock()
	delete(s.element, key)
	s.lock.Unlock()
}
