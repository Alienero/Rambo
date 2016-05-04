package ddl

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/mysql/client"
	"github.com/Alienero/Rambo/rpc"
	"github.com/Alienero/Rambo/util/uuid"

	"github.com/golang/glog"
)

const (
	// CreatDB is the type of the create database plan
	CreatDB = "CreatDB"
)

// DDL is responsible for schema change.
type DDL interface {
	CreateDatabase()
	CreateTable()
	DropTable()
	DropDatabase()
}

// BasePlan is base ddl plan
type BasePlan struct {
	SubPlans    []*SubPlan `json:"sub-plans"`
	LockVersion int64      `json:"lock-version"`
	ID          string     `json:"id"` // uuid
	LockKey     string     `json:"lock-key"`
}

// SubPlan is one of ddl's sub plans
type SubPlan struct {
	Node *meta.MysqlNode `json:"node"`
	SQL  string          `json:"sql"`
}

// CreateDBPlan ddl createDB plan
type CreateDBPlan struct {
	BasePlan
	DBName      string            `json:"db-name"`
	UserName    string            `json:"user-name"`
	FinishNodes []*meta.MysqlNode `json:"finish-nodes"`
}

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
	Seq  string
	User string

	c chan *Result
}

func (t *Task) newC() {
	t.c = make(chan *Result, 1)
}

func (t *Task) done(r *Result) {
	t.c <- r
}

func (t *Task) wait() *Result {
	return <-t.c
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

// NewManage get a new Manage instance
func NewManage(machines []string, localAddr string) *Manage {
	m := &Manage{
		taskQueue: make(chan *Task, 1000),
		w:         NewWait(),
		localAddr: localAddr,
		rs:        rpc.NewGobServer(localAddr),
	}
	if err := m.rs.Register(m); err != nil {
		panic(err)
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
	// start rpc server
	d.rs.ListenAndServe()
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
	// build the ddl plan
	nodes, err := d.info.GetMysqlNodes()
	if err != nil {
		return err
	}
	// get users backend
	if num == 0 {
		num = len(nodes)
	}
	createDB := fmt.Sprintf("CREATE DATABASE %s", path.Join(uname, database))
	subs := make([]*SubPlan, 0, num)
	for i := 0; i < num; i++ {
		index := num % len(nodes)
		node := nodes[index]
		subs = append(subs, &SubPlan{
			SQL:  createDB,
			Node: node,
		})
	}
	var t = &Task{
		Plan: &CreateDBPlan{
			DBName:   database,
			UserName: uname,
			BasePlan: BasePlan{
				SubPlans:    subs,
				LockKey:     path.Join(uname, database),
				ID:          uuid.Get(),
				LockVersion: 0,
			},
		},
		User: uname,
		Type: CreatDB,
		c:    make(chan *Result, 1),
	}
	glog.Infof("DDL: create database task:%+v", t)

	// let master node save task, add etcd queue
	remote, err := d.getMaster(uname)
	if err != nil {
		return err
	}
	// gRPC to master node
	if remote == d.localAddr {
		glog.Infof("using local handle task(%v) plan(%v)", t.Seq, t.Plan.(*CreateDBPlan).ID)
		t.newC()
		d.taskQueue <- t
		r := t.wait()
		return r.err
	}
	// rpc send task
	c, err := rpc.NewGobClient(remote)
	if err != nil {
		return err
	}
	r := new(Result)
	err = c.Call("Manage.SendTask", t, r)
	if err != nil {
		return err
	}
	return r.err
}

// CreateTable get a CreateTable task
// IMPORTANT: only support hash type yet!!!
func (d *Manage) CreateTable() {

}

// DropTable get a DropTable task
func (d *Manage) DropTable() {

}

// DropDatabase get a DropDatabase task
func (d *Manage) DropDatabase() {

}

// SendTask send task to remote server
func (d *Manage) SendTask(t *Task, result *Result) {
	glog.Infof("get a task:%v", t)
	// record plan
	seq, err := d.info.SaveDDLTask(t, t.User)
	if err != nil {
		result.err = err
		return
	}
	t.Seq = seq
	d.taskQueue <- t
}

// CallLock will lock key specific source
func (d *Manage) CallLock(l *DDLLock, result *Result) {
	d.w.Wait(l)
}

// CallUnLock will lock key specific source
func (d *Manage) CallUnLock(l *DDLLock, result *Result) {
	d.w.Continue(l)
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
		t := &Task{
			Seq: node.Key,
			c  : make(chan *Result, 1),
		}
		if err := json.Unmarshal([]byte(node.Value), t); err != nil {
			glog.Infof("Get master(%v)'s tasks error:%v", user, err)
			return
		}
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
		glog.Infof("Manage handle task:%v", t)
		r := d.doTask(t)
		t.done(r)
		if r.err != nil {
			// TODO rollback
			glog.Warningf("DDL execute error:%v", r.err)
		}
	}
}

func (d *Manage) doTask(t *Task) *Result {
	// TODO: impl
	r := new(Result)
	switch t.Type {
	case CreatDB:
		plan := t.Plan.(*CreateDBPlan)
		// check the db
		isExist, err := d.info.IsDBExist(plan.UserName, plan.DBName)
		if err != nil {
			r.err = err
			break
		}
		if isExist {
			r.err = fmt.Errorf("database: %v is already exist in user(%v)", plan.DBName, plan.UserName)
			break
		}
		// execute ddl plan
		// etcd lock
		if err := d.info.Lock(plan.LockKey, plan.ID, plan.LockVersion); err != nil {
			panic(err) // if err is not nil, this node is unreliable
		}
		// TODO:
		// rpc boardcast lock
		plan.LockVersion++
		// get all nodes
		rpcs, err := d.info.GetAllProxyNodes()
		if err != nil {
			// TODO: retry, it must success
			panic(err)
		}
		for _, addr := range rpcs {
			c, err := rpc.NewGobClient(addr)
			if err != nil {
				// TODO: retry, it must success
				panic(err)
			}
			lock := getDDLLock(plan.LockKey, plan.LockVersion, plan.ID)
			if err = c.Call("Manage.CallLock", lock, nil); err != nil {
				// TODO: retry, it must success
				panic(err)
			}
		}
		// dotask
		for _, sp := range plan.SubPlans {
			db, err := client.Open(sp.Node.Host, sp.Node.UserName, sp.Node.Password, "", 0)
			if err != nil {
				// TODO: retry, it must success
				panic(err)
			}
			conn, err := db.GetConn()
			if err != nil {
				// TODO: retry, it must success
				panic(err)
			}
			defer conn.Close()
			_, err = conn.Execute(sp.SQL)
			if err != nil {
				r.err = err
				break
			}
			db.Close()
			// record
			if len(plan.SubPlans) > 1 {
				plan.SubPlans = plan.SubPlans[1:]
				t.Plan = plan
				data, _ := json.Marshal(t)
				plan.FinishNodes = append(plan.FinishNodes, sp.Node)
				d.info.UpdateTask(t.User, t.Seq, string(data))
			} else {
				// register result
				backends := make([]*meta.Backend, 0, len(plan.FinishNodes))
				for _, node := range plan.FinishNodes {
					backend := &meta.Backend{
						Host:       node.Host,
						Name:       node.Name,
						UserName:   node.Password,
						Password:   node.Password,
						ParentNode: node,
					}
					backends = append(backends, backend)
				}
				err = d.info.SaveCreateDatabase(plan.UserName, plan.DBName, t.Seq, backends)
				if err != nil {
					r.err = err
					break
				}
			}
		}

		// unlock etcd lock
		plan.LockVersion++
		if err = d.info.UnLock(plan.LockKey, plan.ID, plan.LockVersion); err != nil {
			// TODO
			panic(err)
		}
		// rpc boardcast unlock
		plan.LockVersion++

		// get all nodes
		rpcs, err = d.info.GetAllProxyNodes()
		if err != nil {
			// TODO: retry, it must success
			panic(err)
		}
		for _, addr := range rpcs {
			c, err := rpc.NewGobClient(addr)
			if err != nil {
				// TODO: retry, it must success
				panic(err)
			}
			lock := getDDLLock(plan.LockKey, plan.LockVersion, plan.ID)
			if err = c.Call("Manage.CallUnLock", lock, nil); err != nil {
				// TODO: retry, it must success
				panic(err)
			}
		}

	default:
		r.err = fmt.Errorf("not support task's type:%v", t.Type)
	}
	return r
}

// BecomeMaster is election's callback method
func (d *Manage) BecomeMaster(key string) {
	go d.getTasks(key)
}

// Result is the result of rpc
type Result struct {
	err error
}

// DDLLock is a rambo lock
type DDLLock struct {
	Key     string
	Version int64
	ID      string
	expired int64
	status  string
}

func getDDLLock(key string, ver int64, id string) *DDLLock {
	return &DDLLock{
		Key:     key,
		Version: ver,
		ID:      id,
	}
}

const (
	StatusLock   = "lock"
	StatusUnlock = "unlock"
)
