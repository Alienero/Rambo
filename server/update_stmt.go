package server

import "github.com/Alienero/Rambo/mysql/sqlparser"

func (sei *session) handleUpdate(stmt *sqlparser.Update) error {
	return sei.handleNormalExecute(stmt)
}
