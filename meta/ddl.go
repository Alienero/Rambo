package meta

import (
	"encoding/json"
	"path"
	"strconv"

	"github.com/coreos/go-etcd/etcd"
)

func (d *Info) Lock(key string, id string, seq int64) error {
	_, err := d.Set(d.lockPath(key), path.Join(id, strconv.FormatInt(seq, 10)), 0)
	return err
}

func (d *Info) UnLock(key string, id string, seq int64) error {
	_, err := d.Update(d.lockPath(key), path.Join(id, strconv.FormatInt(seq, 10)), 0)
	if e, ok := err.(etcd.EtcdError); ok {
		if e.ErrorCode == NotFound {
			err = nil
		}
	}
	return err
}

func (d *Info) lockPath(key string) string {
	return path.Join(Lock, key, "status")
}

// key is user id
func (d *Info) GetMaster(key string) (master string, err error) {
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

// SaveDDLTask save ddl task, return the task id
func (d *Info) SaveDDLTask(v interface{}, user string) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	resp, err := d.CreateInOrder(path.Join(DDLInfo, TaskQueue, user), string(data), 0)
	if err != nil {
		return "", err
	}
	return resp.Node.Key, nil
}

// GetTasks get the user's all tasks
func (d *Info) GetTasks(user string) (etcd.Nodes, error) {
	resp, err := d.Get(path.Join(DDLInfo, TaskQueue, user), true, true)
	if err != nil {
		return nil, err
	}
	return resp.Node.Nodes, nil
}
