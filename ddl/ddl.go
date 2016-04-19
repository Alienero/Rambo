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
type Plan interface {
	Plan()
}

type BasePlan struct {
	SubPlans    []*SubPlan `json:"sub-plans"`
	LockVersion int64      `json:"lock-version"`
	ID          string     `json:"id"` // uuid
	LockKey     string     `json:"lock-key"`
}

type SubPlan struct {
	Node *meta.MysqlNode `json:"node"`
	SQL  string          `json:"sql"`
}

type CreateDBPlan struct {
	BasePlan
	DBName   string `json:"db-name"`
	UserName string `json:"user-name"`
	Password string `json:"password"`
}

func (*CreateDBPlan) Plan() {}

type CreateTablePlan struct {
}

func (*CreateTablePlan) Plan() {}

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
	// fix username and database name
	uname += "/" + database
	database = uname

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
	password := rand.String(15)
	creataUser := fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s'", uname, password)
	createDB := fmt.Sprintf("CREATE DATABASE %s", database)
	grant := fmt.Sprintf("GRANT privileges ON %s.* TO '%s'@'%%'", database, uname)
	subs := make([]*SubPlan, 0, num*3)
	for i := 0; i < num; i++ {
		index := num % len(nodes)
		node := nodes[index]
		subs = append(subs, &SubPlan{
			SQL:  creataUser,
			Node: node,
		})
		subs = append(subs, &SubPlan{
			SQL:  createDB,
			Node: node,
		})
		subs = append(subs, &SubPlan{
			SQL:  grant,
			Node: node,
		})
	}
	var plan Plan = &CreateDBPlan{
		DBName:   database,
		UserName: uname,
		Password: password,
		BasePlan: BasePlan{
			SubPlans:    subs,
			LockKey:     database,
			ID:          uuid.Get(),
			LockVersion: 0,
		},
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
