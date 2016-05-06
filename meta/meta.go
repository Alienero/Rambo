package meta

import (
	"bytes"
	"encoding/json"
	"path"

	"github.com/Alienero/Rambo/mysql"

	"github.com/coreos/go-etcd/etcd"
	"github.com/golang/glog"
)

// Info is metaDB's global public instance.
type Info struct {
	*etcd.Client
}

// NewInfo get a new meta information manage
func NewInfo(machines []string) *Info {
	return &Info{
		Client: etcd.NewClient(machines),
	}
}

// GetAllProxyNodes will get all proxy nodes
func (m *Info) GetAllProxyNodes() ([]string, error) {
	resp, err := m.Get(ProxyNodes, false, true)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		result = append(result, node.Value)
	}
	return result, nil
}

// GetMysqlNodes get all mysql backend nodes
func (m *Info) GetMysqlNodes() ([]*MysqlNode, error) {
	resp, err := m.Get(MysqlInfo, true, true)
	if err != nil {
		return nil, err
	}
	mns := make([]*MysqlNode, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		mn := new(MysqlNode)
		if err = json.Unmarshal([]byte(node.Value), mn); err != nil {
			return nil, err
		}
		mns = append(mns, mn)
	}
	return mns, nil
}

func (m *Info) AddUser(user, password string) error {
	_, err := m.Create(path.Join(UserInfo, user, Password), password, 0)
	return err
}

func (m *Info) GetUserInfo(user string) (*etcd.Response, error) {
	return m.Get(path.Join(UserInfo, user), false, true)
}

func (m *Info) AddBDatabase(db *Database) error {
	return nil
}

func (m *Info) IsDBExist(user, db string) (bool, error) {
	_, err := m.Get(path.Join(UserInfo, user, DB, db), false, false)
	if err != nil {
		if e, ok := err.(*etcd.EtcdError); ok {
			if e.ErrorCode == NotFound {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func (m *Info) CheckUser(user string, auth []byte, salt []byte, db string) bool {
	password, err := m.getPassword(user)
	if err != nil {
		return false
	}

	if isExist := bytes.Equal(auth, mysql.CalcPassword(salt, []byte(password))); !isExist {
		glog.V(2).Infof("Password(!= %v) is nq", password)
		return false
	}
	if db != "" {
		_, err = m.Get(UserInfo+"/"+user+DB+db, false, false)
		if err != nil {
			glog.Infof("Can not get Etcd user db error:%v", err)
			return false
		}

	}
	return true
}

func (m *Info) CheckUserDirect(user, password string) bool {
	p, err := m.getPassword(user)
	if err != nil {
		return false
	}
	return password == p
}

func (m *Info) getPassword(user string) (string, error) {
	resp, err := m.Get(UserInfo+"/"+user+Password, false, false)
	if err != nil {
		glog.Warningf("Get Etcd User password error:%v", err)
		return "", err
	}
	return resp.Node.Value, nil
}

func (m *Info) ShowDatabases(user string) ([]string, error) {
	resp, err := m.Get(path.Join(UserInfo, user, DB), true, true)
	if err != nil {
		if e, ok := err.(*etcd.EtcdError); ok {
			if e.ErrorCode == NotFound {
				return nil, nil
			}
		}
		return nil, err
	}
	s := make([]string, 0, len(resp.Node.Nodes))
	for _, db := range resp.Node.Nodes {
		s = append(s, db.Value)
	}
	return s, nil
}

func (m *Info) ShowTables(user, db string) ([]string, error) {
	return nil, nil
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

// type DBInfo struct {
// 	Tables   []*Table   `json:"tables"`
// 	Backends []*Backend `json:"backends"`
// }

// Table is mysql table meta info.
type Table struct {
	Name         string `json:"name"`
	Scheme       string `json:"scheme"` // default is `hash`
	PartitionKey Key    `json:"partition-key"`
	AutoKeys     []Key  `json:"auto-keys"`
}

type Key struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Length int    `json:"length"`
}

// Backend is one of user's backends
type Backend struct {
	Host       string     `json:"host"`
	UserName   string     `json:"user-name"`
	Password   string     `json:"password"`
	Name       string     `json:"name"` // backend's name
	ParentNode *MysqlNode `json:"parent"`
}

// MysqlNode is one mysql server.
type MysqlNode struct {
	Host     string `json:"host"`
	UserName string `json:"user-name"`
	Password string `json:"password"`
	Name     string `json:"name"` // node's name
}

// Database is global database config.
type Database struct {
	Name     string     `json:"name"`
	Backends []*Backend `json:"backends"`
}
