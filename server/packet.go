package server

import (
	"github.com/Alienero/Rambo/mysql"
)

func (sei *session) writeOK(r *mysql.Result) error {
	if r == nil {
		r = &mysql.Result{Status: sei.status}
	}
	data := make([]byte, 4, 32)

	data = append(data, mysql.OK_HEADER)

	data = append(data, mysql.PutLengthEncodedInt(r.AffectedRows)...)
	data = append(data, mysql.PutLengthEncodedInt(r.InsertId)...)

	if sei.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	return sei.pkg.WritePacket(data)
}

func (sei *session) writeError(e error) error {
	var m *mysql.SqlError
	var ok bool
	if m, ok = e.(*mysql.SqlError); !ok {
		m = mysql.NewError(mysql.ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, mysql.ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if sei.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	return sei.pkg.WritePacket(data)
}

func (sei *session) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, mysql.EOF_HEADER)
	if sei.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return sei.pkg.WritePacket(data)
}

func (sei *session) writeResult(status uint16, r *mysql.Resultset) error {
	columnLen := mysql.PutLengthEncodedInt(uint64(len(r.Fields)))
	data := make([]byte, 4, 1024)
	data = append(data, columnLen...)
	if err := sei.pkg.WritePacket(data); err != nil {
		return err
	}

	for _, v := range r.Fields {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := sei.pkg.WritePacket(data); err != nil {
			return err
		}
	}

	if err := sei.writeEOF(status); err != nil {
		return err
	}

	for _, v := range r.RowDatas {
		data = data[0:4]
		data = append(data, v...)
		if err := sei.pkg.WritePacket(data); err != nil {
			return err
		}
	}

	if err := sei.writeEOF(status); err != nil {
		return err
	}

	return nil
}
