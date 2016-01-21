package meta

import (
	"bytes"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/mysql"

	"github.com/coreos/go-etcd/etcd"
	"github.com/golang/glog"
)

var Meta metaDB

type metaDB struct {
	client *etcd.Client
}

func (m *metaDB) CheckUser(user string, auth []byte, salt []byte, db string) bool {
	resp, err := m.client.Get(UserInfo+"/"+user+Password, false, false)
	if err != nil {
		glog.Warningf("Get Etcd User info error:%v", err)
		return false
	}
	if isExist := bytes.Equal(auth, mysql.CalcPassword(salt, []byte(resp.Node.Value))); !isExist {
		glog.V(2).Infof("Password(!= %v) is nq", resp.Node.Value)
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

// scheme include : special fied, shard scheme, shard key
// TODO: we should cache the information.
func (m *metaDB) GetSchemeTables(user string, db string, table string) (string, []string, error) {
	prefix := UserInfo + "/" + user + DB + db + Tables + table
	scheme, err := m.client.Get(prefix+Scheme, false, false)
	if err != nil {
		glog.Warningf("Get Etcd User table scheme info error:%v", err)
		return "", nil, err
	}
	tables, err := m.client.Get(prefix+ChildTable, true, false)
	if err != nil {
		glog.Warningf("Get Etcd User table's map info error:%v", err)
		return "", nil, err
	}
	ts := make([]string, 0, len(tables.Node.Nodes))
	for _, node := range tables.Node.Nodes {
		ts = append(ts, node.Value)
	}
	return scheme.Node.Value, ts, nil
}

func InitMetaDB() {
	Meta.client = etcd.NewClient(config.Config.Etcd.EtcdAddr)
}

type Scheme struct {
	ScaleType string     `json:"scale-type"` // default is `hash`
	Tables    []string   `json:"tables"`
	Backends  []*Backend `json:"backends"`
}

type Backend struct {
	Host     string `json:"host"`
	UserName string `json:"user-name"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
