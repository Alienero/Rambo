package server

import "github.com/Alienero/Rambo/mysql/sqlparser"

func (sei *session) handleSelect(stmt *sqlparser.Select) error {
	if len(stmt.GroupBy) == 0 {
		// without group by
	}
	return nil
}

func (sei *session) handleSimpleSelect(stmt *sqlparser.Select) error {
	return nil
}
