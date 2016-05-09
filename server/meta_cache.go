package server

import (
	"path"
	"sync"

	"github.com/Alienero/Rambo/meta"

	"github.com/golang/glog"
)

type MetaCache struct {
	cache map[string]*Table
	info  *meta.Info
	sync.RWMutex
}

func NewMetaCache(info *meta.Info) *MetaCache {
	return &MetaCache{
		cache: make(map[string]*Table),
		info:  info,
	}
}

func (m *MetaCache) GetTable(name, db, table string) (*Table, error) {
	m.RLock()
	t, ok := m.cache[path.Join(name, db, table)]
	m.RUnlock()
	if !ok {
		backends, err := m.info.GetDBs(name, db)
		if err != nil {
			return nil, err
		}
		tableInfo, err := m.info.GetTable(name, db, table)
		if err != nil {
			return nil, err
		}
		t = &Table{
			Backends: backends,
			Table:    tableInfo,
		}
		m.Lock()
		m.cache[path.Join(name, db, table)] = t
		m.Unlock()
		glog.Infof("Get Talbe(%s ,db: %s) direct", table, db)
	}
	glog.Infof("Get Talbe(%s ,db: %s) from cache", table, db)
	return t, nil
}

func (m *MetaCache) Del(name string, args ...string) {
	// TODO: impl it
}

type Table struct {
	Backends []*meta.Backend
	Table    *meta.Table
}
