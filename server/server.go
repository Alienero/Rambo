package server

import (
	"net"
	"time"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/util/safemap"

	"github.com/golang/glog"
)

type Server struct {
	addr     string
	seq      uint32
	sessions *safemap.Map
	info     *meta.Info
}

func NewSever() *Server {
	return &Server{
		sessions: safemap.NewMap(),
		addr:     config.Config.Server.ListenAddr,
		info:     meta.NewInfo(config.Config.Etcd.EtcdAddr),
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
		s.sessions.Set(s.seq, sei)
		go sei.serve()
	}
}
