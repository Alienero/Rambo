package server

import (
	"path"
	"strconv"
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
			user:     name,
			db:       db,
			autoKey:  make(map[string]*AutoKey),
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
	autoKey  map[string]*AutoKey

	user string
	db   string
}

type AutoKey struct {
	id  uint64 // last id
	end uint64
}

func (a *AutoKey) OK() bool {
	return a.id < a.end
}

func (a *AutoKey) Get() uint64 {
	a.id++
	return a.id
}

func (a *AutoKey) Set(id, interval uint64) {
	a.id = id
	a.end = a.id + interval
}

func (t *Table) GetKey(field string, interval uint64, info *meta.Info) (string, error) {
	id, ok := t.autoKey[field]
	if ok {
		if id.OK() {
			return strconv.FormatUint(id.Get(), 10), nil
		}
	} else {
		// init
		id = new(AutoKey)
		t.autoKey[field] = id
	}
	start, err := info.GetAutoKey(t.user, t.db, t.Table.Name, field, interval)
	if err != nil {
		return "", err
	}
	id.Set(start, interval)
	return strconv.FormatUint(id.Get(), 10), nil
}
