package server

import (
	"strconv"

	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func (sei *session) handleSet(stmt *sqlparser.Set) error {
	// get first arg
	expr := stmt.Exprs[0]
	name := string(expr.Name.Name)
	switch name {
	case "dbnum":
		buf := sqlparser.NewTrackedBuffer(nil)
		expr.Expr.Format(buf)
		_, ok := expr.Expr.(sqlparser.NumVal)
		if ok {
			i, err := strconv.Atoi(buf.String())
			if err != nil {
				return sei.writeError(mysql.NewDefaultError(mysql.ER_WRONG_VALUE_FOR_VAR, buf.String(), name))
			}
			sei.dbnum = i
			return sei.writeOK(nil)
		}
		return sei.writeError(mysql.NewDefaultError(mysql.ER_WRONG_VALUE_FOR_VAR, buf.String(), name))
	default:
		return sei.writeError(mysql.NewDefaultError(mysql.ER_UNKNOWN_SYSTEM_VARIABLE, name))
	}
}
