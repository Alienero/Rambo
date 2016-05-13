package server

import (
	"strings"

	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func (sei *session) handleExplain(stmt *sqlparser.Explain) error {
	if sei.db == "" {
		return sei.writeError(mysql.NewDefaultError(mysql.ER_NO_DB_ERROR))
	}
	// build plan
	plan, err := sei.buildPlan(stmt.SQL)
	if err != nil {
		return sei.writeError(err)
	}
	// build result
	names := []string{"ID", "Backends", "SQLs"}
	var r *mysql.Resultset
	if len(plan.SQLs) == 0 {
		r, _ = sei.buildEmptySet(names, []interface{}{""})
	} else {
		var values [][]interface{}
		for id, sql := range plan.SQLs {
			values = append(values, []interface{}{
				id,
				sql.Backend.Name,
				strings.Join(sql.SQL, ";"),
			})
		}
		r, _ = sei.buildResultset(nil, names, values)
	}
	return sei.writeResultset(sei.status, r)
}
