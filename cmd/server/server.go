package main

import (
	"flag"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/server"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	// default mysql port.
	config.Config.Server.ListenAddr = "localhost:3306"
	s := server.NewSever()
	// listenning
	s.Run()
}
