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
		glog.Infof("parse sql(%s) error:%v", sql, err)
		return err
	}
	switch v := stmt.(type) {
	case *sqlparser.Explain:

	case *sqlparser.Select:
		if (v.From) == nil {
			switch expr := v.SelectExprs[0].(type) {
			case *sqlparser.NonStarExpr:
				switch f := expr.Expr.(type) {
				case *sqlparser.FuncExpr:
					switch f.Name {
					case "database":
						if f.Exprs == nil {
							// get now database
							name := []string{"DATABASE()"}
							var values [][]interface{}
							if sei.db == "" {
								values = [][]interface{}{
									[]interface{}{"NULL"},
								}
							} else {
								values = [][]interface{}{
									[]interface{}{"NULL"},
								}
							}
							r, err := sei.buildResultset(nil, name, values)
							if err != nil {
								return err
							}
							return sei.writeResultset(sei.status, r)
						}
					}
				}
			case *sqlparser.StarExpr:

			}
		}

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
			id, err := sei.ddlManage().CreateDatabase(sei.user, dbname, 0)
			glog.Infof("DDL plan id(%v)", id)
			return err
		default:
			return fmt.Errorf("create statement %T not support now", stmt)
		}

	case *sqlparser.Show:
		r, err := sei.handleShow(v)
		if err != nil {
			sei.writeEOF(sei.status)
			return err
		}
		return sei.writeResultset(sei.status, r)

	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}
	return nil
}
