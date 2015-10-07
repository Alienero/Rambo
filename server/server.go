package server

import (
	"net"
	"time"

	"github.com/Alienero/Rambo/config"

	"github.com/golang/glog"
)

type Server struct {
	addr string
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		glog.Fatal(err)
	}
	defer l.Close()
	// listen client connection.
	var tempDelay time.Duration
	for {
		rw, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				glog.Warningf("proxy server: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		sei, err := newSession(rw)
		if err != nil {
			continue
		}
		go sei.serve()
	}
}
