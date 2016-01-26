package admin

import (
	"github.com/Alienero/Rambo/meta"
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
func (m *Admin) AddDatabase(db string, n int) {}

// AddTable add a table of a database
func (m *Admin) AddTable() {}

// Check the user
func (m *Admin) checkUser(user, password string) bool {
	m.isLogin = meta.Meta.CheckUserDirect(user, password)
	return m.isLogin
}
