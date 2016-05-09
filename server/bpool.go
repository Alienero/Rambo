package server

import (
	"sync"

	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/mysql/client"
	"github.com/golang/glog"
)

type Bpool struct {
	backends map[string]*client.DB
	sync.RWMutex
}

func NewBpool() *Bpool {
	return &Bpool{
		backends: make(map[string]*client.DB),
	}
}

func (p *Bpool) GetConn(backend *meta.Backend) (*client.Conn, error) {
	// check cahce
	p.RLock()
	db, ok := p.backends[backend.Name]
	p.RUnlock()
	if ok {
		glog.Infof("Get DB(%v) from cache", backend.Name)
		conn, err := db.PopConn()
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
	// create a new one
	var err error
	glog.Infof("Get DB(%v) direct", backend.Name)
	db, err = client.Open(backend.Host, backend.UserName, backend.Password, backend.Name, 0)
	if err != nil {
		return nil, err
	}
	p.Lock()
	srcDB, ok := p.backends[backend.Name]
	if ok {
		db.Close()
		db = srcDB
	}
	p.backends[backend.Name] = db
	p.Unlock()
	return db.PopConn()
}

func (p *Bpool) PushConn(backend *meta.Backend, conn *client.Conn, err error) {
	p.RLock()
	db, ok := p.backends[backend.Name]
	if ok {
		db.PushConn(conn, err)
	} else {
		conn.Close()
	}
	p.RUnlock()
}

func (p *Bpool) CloseDB(backend *meta.Backend) error {
	// impl it
	return nil
}
