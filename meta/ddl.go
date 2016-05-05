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
	return path.Join(DDLInfo, Lock, key, "status")
}

// key is user id
func (d *Info) GetMaster(key string) (master string, err error) {
	resp, err := d.Get(path.Join(DDLInfo, Masters, key), false, false)
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
	resp, err := d.CreateInOrder(path.Join(DDLInfo, user, TaskQueue), string(data), 0)
	if err != nil {
		return "", err
	}
	return resp.Node.Key, nil
}

// GetTasks get the user's all tasks
func (d *Info) GetTasks(user string) (etcd.Nodes, error) {
	resp, err := d.Get(path.Join(DDLInfo, user, TaskQueue), true, true)
	if err != nil {
		return nil, err
	}
	return resp.Node.Nodes, nil
}

// UpdateTask will update this task
func (d *Info) UpdateTask(user, id, value string) error {
	_, err := d.Update(path.Join(DDLInfo, user, TaskQueue, id), value, 0)
	return err
}

// SaveCreateDatabase save db info into etcd
func (d *Info) SaveCreateDatabase(user, db, tid string, backends []*Backend) error {
	for _, backend := range backends {
		data, _ := json.Marshal(backend)
		_, err := d.Set(path.Join(UserInfo, user, DB, db, Backends, backend.Name), string(data), 0)
		if err != nil {
			return err
		}
	}
	_, err := d.Delete(path.Join(DDLInfo, TaskQueue, user, tid), true)
	return err
}

// SetTaskStatus set task's status
func (d *Info) SetTaskStatus(user, id string, data []byte) error {
	_, err := d.Set(path.Join(DDLInfo, user, TaskStatus, id), string(data), 0)
	return err
}

// GetTaskStatus get task's status
func (d *Info) GetTaskStatus(user, id string) ([]byte, error) {
	resp, err := d.Get(path.Join(DDLInfo, user, TaskStatus, id), false, false)
	if err != nil {
		return nil, err
	}
	return []byte(resp.Node.Value), nil
}
