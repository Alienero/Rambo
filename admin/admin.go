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

// Login check the user.
func (m *Admin) Login(user, password string) bool {
	m.isLogin = meta.Meta.CheckUserDirect(user, password)
	return m.isLogin
}
