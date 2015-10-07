package config

import ()

var Config conf

type conf struct {
	Etcd   etcd
	Server server
	Proxy  proxy
}

// Meta infomation.
type etcd struct {
	EtcdAddr string
}

type server struct {
	ListenAddr string
}

type proxy struct{}
