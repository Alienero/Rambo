package server

import (
	"net"
	"runtime"

	"github.com/Alienero/Rambo/mysql"

	"github.com/golang/glog"
)

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

type session struct {
	rw         net.Conn
	id         uint32
	salt       []byte
	status     uint16
	charset    string
	user       string
	db         string
	capability uint32
	pkg        *mysql.PacketIO
	collation  mysql.CollationId
	server     *Server
}

func newSession(rw net.Conn, id uint32, server *Server) (*session, error) {
	sei := &session{
		rw:        rw,
		id:        id,
		salt:      mysql.ScrambleBuf(20),
		pkg:       mysql.NewPacketIO(rw),
		status:    mysql.SERVER_STATUS_AUTOCOMMIT,
		charset:   mysql.DEFAULT_CHARSET,
		collation: mysql.DEFAULT_COLLATION_ID,
		server:    server,
	}
	return sei, nil
}

func (sei *session) serve() {
	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			glog.Error("serve panic %v: %v\n%s", sei.rw.RemoteAddr().String(), err, buf)
		}
		delete(sei.server.sessions, sei.id)
		sei.rw.Close()
	}()
	// handshake.
	if err := sei.writeInitialHandshake(); err != nil {
		glog.Error("server", "Handshake", err.Error(),
			sei.id, "msg", "send initial handshake error")
		return
	}

	if err := sei.readHandshakeResponse(); err != nil {
		glog.Error("server", "readHandshakeResponse",
			err.Error(), sei.id,
			"msg", "read Handshake Response error")

		sei.writeError(err)

		return
	}

	if err := sei.writeOK(nil); err != nil {
		glog.Error("server", "readHandshakeResponse",
			"write ok fail",
			sei.id, "error", err.Error())
		return
	}
}
