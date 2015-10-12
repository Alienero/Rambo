package server

import (
	"net"
	"time"

	"github.com/Alienero/Rambo/config"

	"github.com/golang/glog"
)

type Server struct {
	addr     string
	sessions map[uint32]*session
	seq      uint32
}

func NewSever() *Server {
	return &Server{
		sessions: make(map[uint32]*session),
		addr:     config.Config.Server.ListenAddr,
	}
}

func (s *Server) Run() {
	glog.Info("Server Load will listen on:", s.addr)
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		glog.Fatal(err)
	}
	defer l.Close()
	glog.Info("Server listen on:", s.addr)
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
			glog.Error(e)
			return
		}
		tempDelay = 0
		s.seq++
		sei, err := newSession(rw, s.seq, s)
		if err != nil {
			continue
		}
		s.sessions[s.seq] = sei
		go sei.serve()
	}
}
