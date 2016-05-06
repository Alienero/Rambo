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
	CreateDatabase(uname, database string, num int) (id string, err error)
	CreateTable() (id string, err error)
	DropTable() (id string, err error)
	DropDatabase() (id string, err error)
}

// SubPlan is one of ddl's sub plans
type SubPlan struct {
	Node *meta.MysqlNode `json:"node"`
	SQL  string          `json:"sql"`
}

// Plan is ddl execute plan
type Plan struct {
	SubPlans    []*SubPlan        `json:"sub-plans"`
	LockVersion int64             `json:"lock-version"`
	ID          string            `json:"id"` // uuid
	LockKey     string            `json:"lock-key"`
	DBName      string            `json:"db-name"`
	UserName    string            `json:"user-name"`
	TableName   string            `json:"table-name"`
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
	Plan *Plan
	Type string
	Seq  string

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

// ID get task's id
func (t *Task) ID() string {
	return t.Plan.ID
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
func (d *Manage) CreateDatabase(uname, database string, num int) (string, error) {
	// build the ddl plan
	nodes, err := d.info.GetMysqlNodes()
	if err != nil {
		return "", err
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
		Plan: &Plan{
			DBName:      database,
			UserName:    uname,
			SubPlans:    subs,
			LockKey:     path.Join(uname, database),
			ID:          uuid.Get(),
			LockVersion: 0,
		},
		Type: CreatDB,
		c:    make(chan *Result, 1),
	}
	glog.Infof("DDL: create database task:%+v", t)
	return t.Plan.ID, d.handleTask(t)
}

func (d *Manage) handleTask(t *Task) error {
	// let master node save task, add etcd queue
	remote, err := d.getMaster(t.Plan.UserName)
	if err != nil {
		return err
	}
	// rpc to master node
	if remote == d.localAddr {
		glog.Infof("using local handle task(%v) plan(%v)", t.Seq, t.Plan.ID)
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
func (d *Manage) CreateTable() (id string, err error) {
	return
}

// DropTable get a DropTable task
func (d *Manage) DropTable() (id string, err error) {
	return
}

// DropDatabase get a DropDatabase task
func (d *Manage) DropDatabase() (id string, err error) {
	return
}

// SendTask send task to remote server
func (d *Manage) SendTask(t *Task, result *Result) error {
	glog.Infof("get a task:%v", t)
	// record plan
	seq, err := d.info.SaveDDLTask(t, t.Plan.UserName)
	if err != nil {
		result.err = err
		return err
	}
	status := NewTasksStatus()
	// set up task status monitor
	for _, sp := range t.Plan.SubPlans {
		status.update(sp.Node.Name, Pending, "")
	}
	d.setTaskStatus(t.Plan.UserName, t.ID(), status)
	if err != nil {
		result.err = err
		return err
	}
	t.Seq = seq
	d.taskQueue <- t
	*result = *t.wait()
	return nil
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
			c:   make(chan *Result, 1),
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
	plan := t.Plan
	// get status
	status, err := d.GetTaskStatus(plan.UserName, plan.ID)
	if err != nil {
		r.err = err
		return r
	}
	// prepare handle task
	switch t.Type {
	case CreatDB:
		// check the db
		isExist, err := d.info.IsDBExist(plan.UserName, plan.DBName)
		if err != nil {
			r.err = err
			return r
		}
		if isExist {
			r.err = fmt.Errorf("database: %v is already exist in user(%v)", plan.DBName, plan.UserName)
			return r
		}
	default:
		r.err = fmt.Errorf("not support task's type:%v", t.Type)
	}
	// execute task
	for _, sp := range plan.SubPlans {
		d.updateTaskStatus(plan.UserName, plan.ID, sp.Node.Name, Doing, "", status)
		if err := d.executeSubPlan(sp, t); err != nil {
			d.updateTaskStatus(plan.UserName, plan.ID, sp.Node.Name, Fail, err.Error(), status)
			r.err = err
			break
		} else {
			d.updateTaskStatus(plan.UserName, plan.ID, sp.Node.Name, Done, "", status)
		}
		// record
		if len(plan.SubPlans) > 1 {
			plan.SubPlans = plan.SubPlans[1:]
			data, _ := json.Marshal(t)
			plan.FinishNodes = append(plan.FinishNodes, sp.Node)
			d.info.UpdateTask(t.Plan.UserName, t.Seq, string(data))
		} else {
			// register result
			if err = d.registerTaskResult(t); err != nil {
				r.err = err
				break
			}
		}
	}
	return r
}

func (d *Manage) executeSubPlan(sp *SubPlan, t *Task) error {
	db, err := client.Open(sp.Node.Host, sp.Node.UserName, sp.Node.Password, "", 0)
	if err != nil {
		// TODO: retry, it must success
		return err
	}
	conn, err := db.GetConn()
	if err != nil {
		// TODO: retry, it must success
		return err
	}
	defer conn.Close()
	_, err = conn.Execute(sp.SQL)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (d *Manage) registerTaskResult(t *Task) error {
	switch t.Type {
	case CreatDB:
		backends := make([]*meta.Backend, 0, len(t.Plan.FinishNodes))
		for _, node := range t.Plan.FinishNodes {
			backend := &meta.Backend{
				Host:       node.Host,
				Name:       node.Name,
				UserName:   node.Password,
				Password:   node.Password,
				ParentNode: node,
			}
			backends = append(backends, backend)
		}
		return d.info.SaveCreateDatabase(t.Plan.UserName, t.Plan.DBName, t.Seq, backends)

	default:
		return fmt.Errorf("not support task's type:%v", t.Type)
	}
}

// TasksStatus is task's status
type TasksStatus map[string]*TaskStatus

// NewTasksStatus return a new TasksStatus instance
func NewTasksStatus() *TasksStatus {
	t := make(TasksStatus)
	return &t
}

func (ts *TasksStatus) update(node, status, info string) *TasksStatus {
	m := (*map[string]*TaskStatus)(ts)
	if t, ok := (*m)[node]; ok {
		t.Status = status
		t.Info = info
	} else {
		(*m)[node] = &TaskStatus{
			Status: status,
			Info:   info,
		}
	}
	return ts
}

type TaskStatus struct {
	Status string `json:"node-status"`
	Info   string `json:"info"`
}

const (
	Pending = "wait"
	Doing   = "doing"
	Done    = "done"
	Fail    = "fail"
)

func (d *Manage) updateTaskStatus(uname, id, node, status, info string, ts *TasksStatus) error {
	ts.update(node, status, info)
	return d.setTaskStatus(uname, id, ts)
}

func (d *Manage) setTaskStatus(uname, id string, ts *TasksStatus) error {
	data, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	return d.info.SetTaskStatus(uname, id, data)
}

// GetTaskStatus will get task's status
func (d *Manage) GetTaskStatus(uname, id string) (*TasksStatus, error) {
	data, err := d.info.GetTaskStatus(uname, id)
	if err != nil {
		return nil, err
	}
	t := make(TasksStatus)
	err = json.Unmarshal(data, t)
	return &t, err
}

// BecomeMaster is election's callback method
func (d *Manage) BecomeMaster(key string) {
	go d.getTasks(key)
}

// Result is the result of rpc
type Result struct {
	err  error
	info []byte
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

var _ DDL = &Manage{}
