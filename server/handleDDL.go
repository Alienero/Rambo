package server

import (
	"fmt"

	"github.com/Alienero/Rambo/mysql"

	"github.com/golang/glog"
	"github.com/pingcap/tidb/ast"
	"github.com/pingcap/tidb/parser"
)

func (sei *session) handleDDL(sql string) error {
	stmt, err := parser.ParseOneStmt(sql, "", "")
	if err != nil {
		return err
	}
	switch v := stmt.(type) {
	case *ast.CreateDatabaseStmt:
		dbname := v.Name
		id, rows, err := sei.ddlManage().CreateDatabase(sei.user, dbname, sei.dbnum)
		glog.Infof("DDL plan id(%v)", id)
		if err != nil {
			glog.Infof("CREATE TABLE has an error(%v)", err)
			err = sei.writeError(err)
		} else {
			// one time only creata a db
			r := &mysql.Result{
				AffectedRows: rows,
			}
			err = sei.writeOK(r)
		}
		return err

	case *ast.CreateTableStmt:
		return nil

	default:
		return fmt.Errorf("create statement %T not support now", stmt)
	}
}
