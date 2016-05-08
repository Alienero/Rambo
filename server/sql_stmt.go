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
		sei.writeError(mysql.NewDefaultError(mysql.ER_SYNTAX_ERROR))
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

	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}
	return nil
}
