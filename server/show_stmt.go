package server

import (
	"strings"

	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"
	"github.com/golang/glog"
)

func (sei *session) handleShow(stmt *sqlparser.Show) (*mysql.Resultset, error) {
	var (
		r      = new(mysql.Resultset)
		name   []string
		values [][]interface{}
		err    error
	)

	switch strings.ToLower(stmt.Key) {
	case "databases":
		name = append(name, "Database")
		// get databases from etcd
		var dbs []string
		dbs, err = sei.server.info.ShowDatabases(sei.user)
		if err != nil {
			return nil, err
		}
		if len(dbs) > 0 {
			values = make([][]interface{}, 0, len(dbs))
		} else {
			glog.Info("empty databases")
			r, err = sei.buildEmptySet(name, []interface{}{""})
			break
		}
		for _, db := range dbs {
			values = append(values, []interface{}{db})
		}
		r, err = sei.buildResultset(nil, name, values)
	case "tables":
		if sei.db == "" {
			return nil, mysql.NewError(mysql.ER_NO_DB_ERROR, mysql.MySQLErrName[mysql.ER_NO_DB_ERROR])
		}

	case "ddl_task":
		// get tasks
		// TODO

	}
	return r, err
}
