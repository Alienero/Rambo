package server

import (
	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func (sei *session) handleUseDB(stmt *sqlparser.UseDB) error {
	return sei.useDB(stmt.DB)
}

func (sei *session) useDB(db string) error {
	isExist, err := sei.server.info.IsDBExist(sei.user, db)
	if err != nil {
		return err
	}
	if isExist {
		sei.db = db
		return sei.writeOK(nil)
	}
	return sei.writeError(mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db))
}
