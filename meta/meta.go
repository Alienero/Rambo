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

func (m *metaDB) CheckUser(user, auth string, salt []byte) bool {
	resp, e := m.client.Get(UserInfo+"/"+user+Password, false, false)
	if err := etcdErrCheck(e, NotFoud); err != nil {
		glog.Warningf("Get Etcd User info error:%v", err)
		return false
	}
	return bytes.Equal([]byte(auth), mysql.CalcPassword(salt, []byte(resp.Node.Value)))
}

// scheme include : special fied, shard scheme, shard key
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

type errorCode int

func etcdErrCheck(err error, eq ...errorCode) error {
	if ee, ok := err.(*etcd.EtcdError); ok {
		for _, code := range eq {
			if ee.ErrorCode == int(code) {
				return nil
			}
		}
	}
	return err
}
