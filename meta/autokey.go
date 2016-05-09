package meta

import (
	"path"
	"strconv"

	"github.com/coreos/go-etcd/etcd"
)

func (m *Info) GetAutoKey(user, db, table, key string, interval uint64) (uint64, error) {
	filepath := path.Join(AutoKey, user, db, DB, Tables, table, key)
try:
	resp, err := m.Get(filepath, false, false)
	if err != nil {
		if e, ok := err.(*etcd.EtcdError); ok {
			if e.ErrorCode == NotFound {
				// Create
				if _, err = m.Create(filepath, strconv.FormatUint(interval, 10), 0); err != nil {
					if err.(etcd.EtcdError).ErrorCode == AlreadyExist {
						goto try
					} else {
						return 0, err
					}
				} else {
					return 0, nil
				}
			}
		}
		return 0, err
	}
	last, err := strconv.ParseUint(resp.Node.Value, 10, 64)
	if err != nil {
		return 0, err
	}
	_, err = m.CompareAndSwap(filepath, strconv.FormatUint(last+interval, 10), 0, resp.Node.Value, 0)
	if err != nil {
		if err.(*etcd.EtcdError).ErrorCode == NotEqual {
			goto try
		} else {
			return 0, err
		}
	}
	return last, nil
}
