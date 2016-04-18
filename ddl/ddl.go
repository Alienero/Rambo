package ddl

import (
	"fmt"
	"path"

	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/util/rand"
	"github.com/Alienero/Rambo/util/uuid"

	"github.com/golang/glog"
)

// DDL is responsible for schema change.
type DDL interface {
	CreateDatabase()
	CreateTable()
	DropTable()
	DropDatabase()
}

// Plan is DDL execute plan
type Plan struct {
	SubPlans    []*subPlan `json:"sub-plans"`
	LockVersion int64      `json:"lock-version"`
	ID          string     `json:"id"` // uuid
	LockKey     string     `json:"lock-key"`
}

type subPlan struct {
	Node *meta.MysqlNode `json:"node"`
	SQL  string          `json:"sql"`
}

// TODO:
// 1 get ddl task from etcd, etcd lock
// 2 gRPC get ddl task
// 3 do task
// 4 recorder tasks

// Manage manage handle all ddl stmt
type Manage struct {
	taskQueue chan *Plan
	w         Waitter
	localAddr string

	election *meta.Election
}

func (d *Manage) getMaster(key string) (string, error) {
	// should create a master node
	master, err := d.election.GetMaster(path.Join(meta.Masters, key))
	if err != nil {
		return "", err
	}
	return master, nil
}

// CreateDatabase will create a new database for the uname's user
func (d *Manage) CreateDatabase(uname, database string, num int) error {
	// check the db
	isExist, err := meta.Info.IsDBExist(uname, database)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("database: %v is already exist", database)
	}

	// build the ddl plan
	nodes, err := meta.Info.GetMysqlNodes()
	if err != nil {
		return err
	}
	// get users backend
	if num == 0 {
		num = len(nodes)
	}
	// SQL: CREATE USER 'pig'@'%' IDENTIFIED BY '123456';
	//      GRANT privileges ON databasename.* TO 'username'@'%'
	uname += "/" + database
	database = uname
	password := rand.String(15)
	creataUser := fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s';", uname, password)
	createDB := fmt.Sprintf("CREATE DATABASE %s", database)
	grant := fmt.Sprintf("GRANT privileges ON %s.* TO '%s'@'%%'", database, uname)
	subs := make([]*subPlan, 0, num*3)
	for i := 0; i < num; i++ {
		index := num % len(nodes)
		node := nodes[index]
		subs = append(subs, &subPlan{
			SQL:  creataUser,
			Node: node,
		})
		subs = append(subs, &subPlan{
			SQL:  createDB,
			Node: node,
		})
		subs = append(subs, &subPlan{
			SQL:  grant,
			Node: node,
		})
	}
	plan := &Plan{
		SubPlans:    subs,
		LockKey:     database,
		ID:          uuid.Get(),
		LockVersion: 0,
	}
	glog.Infof("DDL: create database plan:%v", plan)

	// TODO: save plan, add etcd queue
	master, err := d.getMaster(uname)
	if err != nil {
		return err
	}
	// gRPC to master node
	if master == d.localAddr {

	}
	return nil
}

func (d *Manage) buildPlan() {}

// IMPORTANT: only support hash type yet!!!
func (d *Manage) CreateTable() {

}

func (d *Manage) DropTable() {

}

func (d *Manage) DropDatabase() {

}

func (d *Manage) doTask() {

}
