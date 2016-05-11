package server

import "github.com/Alienero/Rambo/mysql/sqlparser"

func (sei *session) handleDelete(stmt *sqlparser.Delete) error {
	return sei.handleNormalExecute(stmt)
}
