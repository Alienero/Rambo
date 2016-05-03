package server

import (
	"fmt"
	"strings"

	"github.com/Alienero/Rambo/mysql/sqlparser"

	"github.com/golang/glog"
	"github.com/pingcap/tidb/ast"
	"github.com/pingcap/tidb/parser"
)

func (sei *session) handleQuery(data []byte) error {
	sql := strings.TrimRight(string(data), ";")
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		glog.Infof("parse sql error:%v", err)
		return err
	}
	switch stmt.(type) {
	case *sqlparser.Select:

	case *sqlparser.Insert:

	case *sqlparser.Update:

	case *sqlparser.Delete:

	case *sqlparser.Set:

	case *sqlparser.DDL:
		stmt, err := parser.ParseOneStmt(sql, "", "")
		if err != nil {
			return err
		}
		switch v := stmt.(type) {
		case *ast.CreateDatabaseStmt:
			dbname := v.Name
			return sei.ddlManage().CreateDatabase(sei.user, dbname, 0)
		default:
			return fmt.Errorf("create statement %T not support now", stmt)
		}

	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}
	return nil
}
