package client

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Alienero/Rambo/mysql"
)

// errors
var (
	ErrDatabaseClose = errors.New("database is close")
	ErrConnIsNil     = errors.New("connection is nil")
	ErrBadConn       = errors.New("connection was bad")
)

const (
	Up = iota
	Down
	ManualDown
	Unknown

	InitConnCount           = 16
	DefaultMaxConnNum       = 1024
	PingPeroid        int64 = 4
)

type DB struct {
	sync.RWMutex

	addr     string
	user     string
	password string
	db       string
	state    int32

	maxConnNum  int
	InitConnNum int
	idleConns   chan *Conn
	cacheConns  chan *Conn
	checkConn   *Conn
}

func Open(addr string, user string, password string, dbName string, maxConnNum int) (*DB, error) {
	var err error
	db := new(DB)
	db.addr = addr
	db.user = user
	db.password = password
	db.db = dbName

	if 0 < maxConnNum {
		db.maxConnNum = maxConnNum
		if db.maxConnNum < 16 {
			db.InitConnNum = db.maxConnNum
		} else {
			db.InitConnNum = db.maxConnNum / 4
		}
	} else {
		db.maxConnNum = DefaultMaxConnNum
		db.InitConnNum = InitConnCount
	}
	//check connection
	db.checkConn, err = db.newConn()
	if err != nil {
		db.Close()
		return nil, ErrDatabaseClose
	}

	db.idleConns = make(chan *Conn, db.maxConnNum)
	db.cacheConns = make(chan *Conn, db.maxConnNum)
	atomic.StoreInt32(&(db.state), Unknown)

	for i := 0; i < db.maxConnNum; i++ {
		if i < db.InitConnNum {
			conn, err := db.newConn()
			if err != nil {
				db.Close()
				return nil, ErrDatabaseClose
			}
			conn.pushTimestamp = time.Now().Unix()
			db.cacheConns <- conn
		} else {
			conn := new(Conn)
			db.idleConns <- conn
		}
	}

	return db, nil
}

func (db *DB) Addr() string {
	return db.addr
}

func (db *DB) State() string {
	var state string
	switch db.state {
	case Up:
		state = "up"
	case Down, ManualDown:
		state = "down"
	case Unknown:
		state = "unknow"
	}
	return state
}

func (db *DB) IdleConnCount() int {
	db.RLock()
	defer db.RUnlock()
	return len(db.cacheConns)
}

func (db *DB) Close() error {
	db.Lock()
	idleChannel := db.idleConns
	cacheChannel := db.cacheConns
	db.cacheConns = nil
	db.idleConns = nil
	db.Unlock()
	if cacheChannel == nil || idleChannel == nil {
		return nil
	}

	close(cacheChannel)
	for conn := range cacheChannel {
		db.closeConn(conn)
	}
	close(idleChannel)

	return nil
}

func (db *DB) getConns() (chan *Conn, chan *Conn) {
	db.RLock()
	cacheConns := db.cacheConns
	idleConns := db.idleConns
	db.RUnlock()
	return cacheConns, idleConns
}

func (db *DB) getCacheConns() chan *Conn {
	db.RLock()
	conns := db.cacheConns
	db.RUnlock()
	return conns
}

func (db *DB) getIdleConns() chan *Conn {
	db.RLock()
	conns := db.idleConns
	db.RUnlock()
	return conns
}

func (db *DB) Ping() error {
	var err error
	if db.checkConn == nil {
		db.checkConn, err = db.newConn()
		if err != nil {
			db.closeConn(db.checkConn)
			db.checkConn = nil
			return err
		}
	}
	err = db.checkConn.Ping()
	if err != nil {
		db.closeConn(db.checkConn)
		db.checkConn = nil
		return err
	}
	return nil
}

func (db *DB) newConn() (*Conn, error) {
	co := new(Conn)

	if err := co.Connect(db.addr, db.user, db.password, db.db); err != nil {
		return nil, err
	}

	return co, nil
}

func (db *DB) closeConn(co *Conn) error {
	if co != nil {
		co.Close()
		conns := db.getIdleConns()
		if conns != nil {
			select {
			case conns <- co:
				return nil
			default:
				return nil
			}
		}
	}
	return nil
}

func (db *DB) tryReuse(co *Conn) error {
	var err error
	//reuse Connection
	if co.IsInTransaction() {
		//we can not reuse a connection in transaction status
		err = co.Rollback()
		if err != nil {
			return err
		}
	}

	if !co.IsAutoCommit() {
		//we can not  reuse a connection not in autocomit
		_, err = co.exec("set autocommit = 1")
		if err != nil {
			return err
		}
	}

	//connection may be set names early
	//we must use default utf8
	if co.GetCharset() != mysql.DEFAULT_CHARSET {
		err = co.SetCharset(mysql.DEFAULT_CHARSET)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) PopConn() (*Conn, error) {
	var co *Conn
	var err error

	cacheConns, idleConns := db.getConns()
	if cacheConns == nil || idleConns == nil {
		return nil, ErrDatabaseClose
	}
	co = db.GetConnFromCache(cacheConns)
	if co == nil {
		co, err = db.GetConnFromIdle(cacheConns, idleConns)
		if err != nil {
			return nil, err
		}
	}

	err = db.tryReuse(co)
	if err != nil {
		db.closeConn(co)
		return nil, err
	}

	return co, nil
}

func (db *DB) GetConnFromCache(cacheConns chan *Conn) *Conn {
	var co *Conn
	var err error
	for 0 < len(cacheConns) {
		co = <-cacheConns
		if co != nil && PingPeroid < time.Now().Unix()-co.pushTimestamp {
			err = co.Ping()
			if err != nil {
				db.closeConn(co)
				co = nil
			}
		}
		if co != nil {
			break
		}
	}
	return co
}

func (db *DB) GetConnFromIdle(cacheConns, idleConns chan *Conn) (*Conn, error) {
	var co *Conn
	var err error
	select {
	case co = <-idleConns:
		err = co.Connect(db.addr, db.user, db.password, db.db)
		if err != nil {
			db.closeConn(co)
			return nil, err
		}
		return co, nil
	case co = <-cacheConns:
		if co == nil {
			return nil, ErrConnIsNil
		}
		if co != nil && PingPeroid < time.Now().Unix()-co.pushTimestamp {
			err = co.Ping()
			if err != nil {
				db.closeConn(co)
				return nil, ErrBadConn
			}
		}
	}
	return co, nil
}

func (db *DB) PushConn(co *Conn, err error) {
	if co == nil {
		return
	}
	conns := db.getCacheConns()
	if conns == nil {
		co.Close()
		return
	}
	if err != nil {
		db.closeConn(co)
		return
	}
	co.pushTimestamp = time.Now().Unix()
	select {
	case conns <- co:
		return
	default:
		db.closeConn(co)
		return
	}
}

type BackendConn struct {
	*Conn
	db *DB
}

func (p *BackendConn) Close() {
	if p != nil && p.Conn != nil {
		if p.Conn.pkgErr != nil {
			p.db.closeConn(p.Conn)
		} else {
			p.db.PushConn(p.Conn, nil)
		}
		p.Conn = nil
	}
}

func (db *DB) GetConn() (*BackendConn, error) {
	c, err := db.PopConn()
	if err != nil {
		return nil, err
	}
	return &BackendConn{c, db}, nil
}
