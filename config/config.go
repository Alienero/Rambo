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
	// mutil etcd nodes.
	EtcdAddr []string
	Salt     []byte
}

type server struct {
	ListenAddr string
}

type proxy struct{}
