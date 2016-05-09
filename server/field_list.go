package server

import (
	"bytes"

	"github.com/Alienero/Rambo/mysql"
)

func (sei *session) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	if sei.db == "" {
		return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
	}

	tableInfo, err := sei.getMeta().GetTable(sei.user, sei.db, table)
	if err != nil {
		return err
	}

	co, err := sei.getBpool().GetConn(tableInfo.Backends[0])
	if err != nil {
		return err
	}
	defer sei.getBpool().PushConn(tableInfo.Backends[0], co, err)
	fs, err := co.FieldList(table, wildcard)
	if err != nil {
		return err
	}
	return sei.writeFieldList(sei.status, fs)
}

func (sei *session) writeFieldList(status uint16, fs []*mysql.Field) error {

	data := make([]byte, 4, 1024)

	for _, v := range fs {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := sei.pkg.WritePacket(data); err != nil {
			return err
		}
	}

	if err := sei.writeEOF(status); err != nil {
		return err
	}
	return nil
}
