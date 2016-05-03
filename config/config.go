package config

var Config conf

type conf struct {
	Etcd   etcd
	Server server
	Proxy  proxy
}

// Meta infomation.
type etcd struct {
	// mutil etcd nodes.
	EtcdAddr  []string
	Salt      []byte
	UpdateTTL uint64
}

type server struct {
	ListenAddr string
	RPCAddr    string
}

type proxy struct {
}
