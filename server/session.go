package server

import (
	"bytes"
	"fmt"
	"net"
	"runtime"

	"github.com/Alienero/Rambo/ddl"
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
	closed     bool   // is session closed.
	password   string // backend psw

	// system's args
	dbnum        int
	partitionKey string
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

func (sei *session) Close() error {
	if sei.closed {
		return nil
	}
	sei.closed = true
	return sei.rw.Close()
}

func (sei *session) serve() {
	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			glog.Errorf("serve panic %v: %v\n%s", sei.rw.RemoteAddr().String(), err, buf)
		}
		sei.server.sessions.Delete(sei.id)
		sei.Close()
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

	sei.pkg.Sequence = 0

	// Read Packet.
	for {
		data, err := sei.pkg.ReadPacket()
		if err != nil {
			glog.Errorf("Proxy read packet error:%v", err)
			return
		}
		if err = sei.dispatch(data); err != nil {
			glog.Errorf("Proxy dispatch com error:%v", err)
			return
		}
		if sei.closed {
			return
		}
		sei.pkg.Sequence = 0
	}
}

func (sei *session) dispatch(data []byte) error {
	cmd := data[0]
	glog.Infof("cmd type: %v", cmd)
	data = data[1:]

	switch cmd {
	case mysql.COM_QUIT:
		return sei.Close()
	case mysql.COM_QUERY:
		// TODO: test stmt,should rm.
		fmt.Println(cmd, string(data))
		return sei.handleQuery(data)
	case mysql.COM_PING:
		return sei.writeOK(nil)
	case mysql.COM_INIT_DB:
		return sei.useDB(string(data))
	case mysql.COM_FIELD_LIST:
		// get the column definitions of a table
		index := bytes.IndexByte(data, 0x00)
		table := string(data[0:index])
		wildcard := string(data[index+1:])
		glog.Info(table, "    ", wildcard)
	case mysql.COM_STMT_PREPARE:
	case mysql.COM_STMT_EXECUTE:
	case mysql.COM_STMT_CLOSE:
	case mysql.COM_STMT_SEND_LONG_DATA:
	case mysql.COM_STMT_RESET:
	case mysql.COM_SET_OPTION:
	default:
		msg := fmt.Sprintf("command %d not supported", cmd)
		glog.Errorf("session(%v) dispatch %s", sei.id, msg)
		return mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}
	return nil
}

func (sei *session) ddlManage() *ddl.Manage {
	return sei.server.ddlManage
}
