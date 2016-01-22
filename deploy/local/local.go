package main

import (
	"encoding/json"
	"flag"
	"log"
	"path"
	"strings"

	"github.com/Alienero/Rambo/meta"

	"github.com/coreos/go-etcd/etcd"
)

var (
	user     = flag.String("user", "root", "-user=root")
	host     = flag.String("host", "localhost:33306", "-host=localhost:3306")
	password = flag.String("password", "password", "-password=password")
	name     = flag.String("name", "db0", "-name=db0")

	etcdHost = flag.String("etcd_host", "http://localhost:4001", "-etcd_host=http://localhost:4001")
)

func main() {
	flag.Parse()

	machines := strings.Split(*etcdHost, ",")
	client := etcd.NewClient(machines)
	backend := &meta.MysqlNode{
		Host:     *host,
		UserName: *user,
		Password: *password,
		Name:     *name,
	}
	data, err := json.Marshal(backend)
	if err != nil {
		log.Fatalf("marshal backend error: %v", err)
	}
	_, err = client.Set(path.Join(meta.MysqlInfo, *name), string(data), 0)
	if err != nil {
		log.Fatalf("set mysql node error: %v", err)
	}
	log.Printf("set mysql node(%v) success!", *name)
}
