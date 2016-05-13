package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var Config conf

type conf struct {
	IsDev      bool
	LogFileDir string

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
	AutoKeyInterval uint64
}

func InitConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	return toml.Unmarshal(data, &Config)
}
