package meta

import (
	"path"
	"strconv"

	"github.com/coreos/go-etcd/etcd"
)

type DDL struct {
	*etcd.Client
}

func NewDDL(machines []string) *DDL {
	return &DDL{
		Client: etcd.NewClient(machines),
	}
}

func (d *DDL) Lock(key string, id string, seq int64) error {
	_, err := d.Set(d.lockPath(key), path.Join(id, strconv.FormatInt(seq, 10)), 0)
	return err
}

func (d *DDL) UnLock(key string, id string, seq int64) error {
	_, err := d.CompareAndDelete(d.lockPath(key), path.Join(id, strconv.FormatInt(seq, 10)), 0)
	if e, ok := err.(etcd.EtcdError); ok {
		if e.ErrorCode == NotFound {
			err = nil
		}
	}
	return err
}

func (d *DDL) lockPath(key string) string {
	return path.Join(Lock, key, "status")
}

// key is user id
func (d *DDL) GetMaster(key string) (master string, err error) {
	resp, err := d.Get(path.Join(Masters, key), false, false)
	if err != nil {
		if e, ok := err.(etcd.EtcdError); ok {
			if e.ErrorCode == NotFound {
				return "", nil
			}
		}
		return "", err
	}
	return resp.Node.Value, nil
}
