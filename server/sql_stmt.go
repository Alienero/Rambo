package server

import (
	"fmt"
	"strings"

	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"

	"github.com/golang/glog"
)

func (sei *session) handleQuery(data []byte) error {
	sql := strings.TrimRight(string(data), ";")
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		glog.Infof("parse sql(%s) error:%v", sql, err)
		return sei.writeError(mysql.NewDefaultError(mysql.ER_SYNTAX_ERROR))
	}
	switch v := stmt.(type) {
	case *sqlparser.Explain:
		return sei.handleExplain(v)

	case *sqlparser.Select:
		return sei.handleSelect(v)

	case *sqlparser.Insert:
		return sei.handleInsert(v)

	case *sqlparser.Update:
		return sei.handleUpdate(v)

	case *sqlparser.Delete:
		return sei.handleDelete(v)

	case *sqlparser.Set:
		// only support like `SET autocommit=1`
		return sei.handleSet(v)

	case *sqlparser.DDL:
		return sei.handleDDL(sql)

	case *sqlparser.Show:
		r, err := sei.handleShow(v)
		if err != nil {
			glog.Infof("handle show stmt has error:%v", err)
			sei.writeError(err)
			// not throw the error
			return nil
		}
		return sei.writeResultset(sei.status, r)

	case *sqlparser.UseDB:
		return sei.handleUseDB(v)

	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}
	return nil
}

func (sei *session) handleNormalExecute(stmt sqlparser.Statement) error {
	if sei.db == "" {
		return sei.writeError(mysql.NewDefaultError(mysql.ER_NO_DB_ERROR))
	}

	plan, err := sei.buildPlan(stmt)
	if err != nil {
		return sei.writeError(err)
	}
	rs, err := sei.executePlan(plan)
	if err != nil {
		return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error()))
	}
	// merge results
	return sei.mergeExecResult(rs)
}
