package server

import (
	"net"

	"github.com/golang/glog"
)

type session struct{}

func newSession(rw net.Conn) *session {
	return nil
}
