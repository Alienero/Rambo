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
	config.Config.Server.ListenAddr = "localhost:3306"
	s := server.NewSever()
	s.Run()
}
