package ddl

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/rpc"
	"github.com/Alienero/Rambo/util/rand"
	"github.com/Alienero/Rambo/util/uuid"

	"github.com/golang/glog"
)

const (
	CreatDB = "CreatDB"
)

// DDL is responsible for schema change.
type DDL interface {
	CreateDatabase()
	CreateTable()
	DropTable()
	DropDatabase()
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

// func (c *CreateDBPlan) Plan() string {
// 	return c.User
// }

// type CreateTablePlan struct {
// }

// func (*CreateTablePlan) Plan() {}

// TODO:
// 1 get ddl task from etcd, etcd lock
// 2 gRPC get ddl task
// 3 do task
// 4 recorder tasks

var errUnknowPlan = errors.New("unknow plan")

// Task is DDL execute plan
type Task struct {
	Plan interface{}
	Type string
	ID   string
	User string
}

// Manage manage handle all ddl stmt
type Manage struct {
	taskQueue chan *Task
	w         Waitter
	localAddr string // rpc addr

	rs rpc.Server

	election *meta.Election
	info     *meta.Info
}

func NewManage(machines []string, localAddr string) *Manage {
	m := &Manage{
		taskQueue: make(chan *Task, 1000),
		w:         NewWait(),
		localAddr: localAddr,
		rs:        rpc.NewGobServer(localAddr),
	}
	m.election = meta.NewElection(machines, path.Join(meta.DDLInfo, meta.Masters),
		uint64(time.Second*20), m.BecomeMaster, localAddr)
	return m
}

// Run bootstrap ddl manage
func (d *Manage) Run() {
	// watch all masters
	go d.election.Watch()
	// handle tasks
	go d.doTasks()
	// TODO: start rpc server
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
	isExist, err := d.info.IsDBExist(uname, database)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("database: %v is already exist", database)
	}
	username := uname
	// fix username and database name
	uname += "/" + database
	database = uname

	// build the ddl plan
	nodes, err := d.info.GetMysqlNodes()
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
	var t = &Task{
		Plan: &CreateDBPlan{
			DBName:   database,
			UserName: uname,
			Password: password,
			BasePlan: BasePlan{
				SubPlans:    subs,
				LockKey:     database,
				ID:          uuid.Get(),
				LockVersion: 0,
			},
		},
		User: username,
		Type: CreatDB,
	}
	glog.Infof("DDL: create database task:%+v", t)

	// TODO: save plan, add etcd queue
	master, err := d.getMaster(username)
	if err != nil {
		return err
	}
	// gRPC to master node
	if master == d.localAddr {
		d.taskQueue <- t
	} else {
		// TODO rpc to remote server
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

// SendTask send task to remote server
func (d *Manage) SendTask(t *Task, result *Result) {
	// record plan
	id, err := d.info.SaveDDLTask(t, t.User)
	if err != nil {
		result.err = err
		return
	}
	t.ID = id
	d.taskQueue <- t
}

func (d *Manage) getTasks(key string) {
	// get master's tasks
	if !strings.HasPrefix(key, path.Join(meta.DDLInfo, meta.TaskQueue)) {
		return
	}
	_, user := path.Split(key)
	nodes, err := d.info.GetTasks(user)
	if err != nil {
		glog.Infof("Get master(%v)'s tasks error:%v", user, err)
		return
	}
	for _, node := range nodes {
		t := new(Task)
		if err := json.Unmarshal([]byte(node.Value), t); err != nil {
			glog.Infof("Get master(%v)'s tasks error:%v", user, err)
			return
		}
		t.ID = node.Key
		d.taskQueue <- t
	}
}

func (d *Manage) doTasks() {
	for {
		t, ok := <-d.taskQueue
		if !ok {
			glog.Info("Manage is closed")
			return
		}
		glog.Info(t)
		// TODO: impl
	}
}

func (d *Manage) doTask(t *Task) {
	// TODO: impl
}

// BecomeMaster is election's callback method
func (d *Manage) BecomeMaster(key string) {
	go d.getTasks(key)
}

type Result struct {
	err error
}
