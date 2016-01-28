package admin

import (
	"errors"

	"github.com/Alienero/Rambo/meta"
)

var (
	errInvalidUser = errors.New("invalid user")
	errDBExisted   = errors.New("db is arlready existed")
)

// Admin provide some method to manage database
type Admin struct {
	isLogin bool
}

// AddUser add a database user
func (Admin) AddUser(user, password string) error {
	return meta.Meta.AddUser(user, password)
}

// AddDatabase a database for user, db is the database name,
// n is create how many partition database will create.
func (m *Admin) AddDatabase(db, user string, n int) error {
	if !m.isLogin {
		return errInvalidUser
	}
	isExist, err := meta.Meta.IsDBExist(db)
	if err != nil {
		return err
	}
	if isExist {
		return errDBExisted
	}
	// create database.
	// get backends.
	return nil
}

// AddTable add a table of a database.
func (m *Admin) AddTable(tableName, dbName string) {}

// Login check the user.
func (m *Admin) Login(user, password string) bool {
	m.isLogin = meta.Meta.CheckUserDirect(user, password)
	return m.isLogin
}
