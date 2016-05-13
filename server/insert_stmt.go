package server

import (
	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"

	"github.com/golang/glog"
)

func (sei *session) handleInsert(stmt *sqlparser.Insert) error {
	return sei.handleNormalExecute(stmt)
}

func (sei *session) mergeExecResult(rs []*mysql.Result) error {
	r := new(mysql.Result)
	for _, v := range rs {
		r.Status |= v.Status
		r.AffectedRows += v.AffectedRows
		if r.InsertId == 0 {
			r.InsertId = v.InsertId
		} else if r.InsertId > v.InsertId {
			// last insert id is first gen id for multi row inserted
			// see http://dev.mysql.com/doc/refman/5.6/en/information-functions.html#function_last-insert-id
			r.InsertId = v.InsertId
		}
	}

	if r.InsertId > 0 {
		sei.lastInsertId = int64(r.InsertId)
	}
	glog.Infof("last insert id:%d", r.InsertId)
	return sei.writeOK(r)
}
