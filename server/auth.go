package server

import (
	"bytes"
	"encoding/binary"

	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/mysql"

	"github.com/golang/glog"
)

// https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
func (sei *session) writeInitialHandshake() error {
	data := make([]byte, 4, 128)
	// [0a] protocol version
	data = append(data, 0x0a)
	// server version
	data = append(data, []byte(mysql.ServerVersion)...)
	data = append(data, 0)
	// connection id
	data = append(data, byte(sei.id), byte(sei.id>>8), byte(sei.id>>16), byte(sei.id>>24))
	// auth-plugin-data-part-1
	data = append(data, sei.salt[0:8]...)
	// filler [00]
	data = append(data, 0)
	// capability flag lower 2 bytes, using default capability here.
	data = append(data, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))
	// charset, utf-8 default
	data = append(data, uint8(mysql.DEFAULT_COLLATION_ID))
	// status
	data = append(data, byte(sei.status), byte(sei.status>>8))
	// below 13 byte may not be used
	// capability flag upper 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY>>16), byte(DEFAULT_CAPABILITY>>24))
	// not enable CLIENT_PLUGIN_AUTH
	data = append(data, 0)
	// reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	// auth-plugin-data-part-2
	data = append(data, sei.salt[8:]...)
	data = append(data, 0)
	return sei.pkg.WritePacket(data)
}

func (sei *session) readHandshakeResponse() error {
	data, err := sei.pkg.ReadPacket()
	if err != nil {
		return err
	}

	pos := 0

	// capability
	sei.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4

	// skip max packet size
	pos += 4

	// charset, skip, if you want to use another charset, use set names
	// sei.collation = CollationId(data[pos])
	pos++

	// skip reserved 23[00]
	pos += 23

	// user name
	sei.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(sei.user) + 1

	// for debug
	glog.Info("User:", sei.user)

	// auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]

	pos += authLen

	if sei.capability&mysql.CLIENT_CONNECT_WITH_DB > 0 {
		if len(data[pos:]) == 0 {
			return nil
		}

		sei.db = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(sei.db) + 1
		glog.Info("DB is :", sei.db)
	}

	// user and password check.
	if !meta.Info.CheckUser(sei.user, auth, sei.salt, sei.db) {
		glog.Infof("User(%v) password or user name error", sei.user)
		return mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR)
	}

	// user db.
	if sei.db != "" {
		// if err := c.useDB(db); err != nil {
		// 	return err
		// }
	}

	return nil
}
