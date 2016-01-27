package meta

import (
	"bytes"
	"path"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/mysql"

	"github.com/coreos/go-etcd/etcd"
	"github.com/golang/glog"
)

// Meta is metaDB's global public instance.
var Meta metaDB

type metaDB struct {
	client *etcd.Client
}

func (m *metaDB) AddUser(user, password string) error {
	_, err := m.client.Create(path.Join(UserInfo, user, Password), password, 0)
	return err
}

func (m *metaDB) AddBDatabase(db *Database) error {
	return nil
}

func (m *metaDB) IsDBExist(db string) (bool, error) {
	_, e := m.client.Get(path.Join(Databases, db), false, false)
	if e != nil {
		if err, ok := e.(*etcd.EtcdError); ok {
			if err.ErrorCode == NotFoud {
				return false, nil
			}
			return false, err
		}
		return false, e
	}
	return true, nil
}

func (m *metaDB) CheckUser(user string, auth []byte, salt []byte, db string) bool {
	password, err := m.getPassword(user)
	if err != nil {
		return false
	}

	if isExist := bytes.Equal(auth, mysql.CalcPassword(salt, []byte(password))); !isExist {
		glog.V(2).Infof("Password(!= %v) is nq", password)
		return false
	}
	if db != "" {
		_, err = m.client.Get(UserInfo+"/"+user+DB+db, false, false)
		if err != nil {
			glog.Infof("Can not get Etcd user db error:%v", err)
			return false
		}

	}
	return true
}

func (m *metaDB) CheckUserDirect(user, password string) bool {
	p, err := m.getPassword(user)
	if err != nil {
		return false
	}
	return password == p
}

func (m *metaDB) getPassword(user string) (string, error) {
	resp, err := m.client.Get(UserInfo+"/"+user+Password, false, false)
	if err != nil {
		glog.Warningf("Get Etcd User password error:%v", err)
		return "", err
	}
	return resp.Node.Value, nil
}

// scheme include : special fied, shard scheme, shard key
// TODO: we should cache the information.
// func (m *metaDB) GetSchemeTables(user string, db string, table string) (string, []string, error) {
// 	prefix := UserInfo + "/" + user + DB + db + Tables + table
// 	scheme, err := m.client.Get(prefix+Scheme, false, false)
// 	if err != nil {
// 		glog.Warningf("Get Etcd User table scheme info error:%v", err)
// 		return "", nil, err
// 	}
// 	tables, err := m.client.Get(prefix+ChildTable, true, false)
// 	if err != nil {
// 		glog.Warningf("Get Etcd User table's map info error:%v", err)
// 		return "", nil, err
// 	}
// 	ts := make([]string, 0, len(tables.Node.Nodes))
// 	for _, node := range tables.Node.Nodes {
// 		ts = append(ts, node.Value)
// 	}
// 	return scheme.Node.Value, ts, nil
// }

// InitMetaDB will Init DB.
func InitMetaDB() {
	Meta.client = etcd.NewClient(config.Config.Etcd.EtcdAddr)
}

type DBInfo struct {
	Scheme string   `json:"scheme"` // default is `hash`
	Tables []*Table `json:"tables"`
}

type Table struct {
	Name         string     `json:"name"`
	PartitionKey string     `json:"partition-key"`
	Backends     []*Backend `json:"backends"`
}

type Backend struct {
	Host     string `json:"host"`
	UserName string `json:"user-name"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// MysqlNode is one mysql server.
type MysqlNode struct {
	Host     string `json:"host"`
	UserName string `json:"user-name"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Database is global database config.
type Database struct {
	Name     string     `json:"name"`
	Backends []*Backend `json:"backends"`
}
