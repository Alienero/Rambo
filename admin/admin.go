package admin

import (
	"github.com/Alienero/Rambo/meta"
)

type Admin struct {
}

func (m *Admin) AddUser(user, password string) error {
	return meta.Meta.AddUser(user, password)
}

func (m *Admin) AddDatabase() {}

func (m *Admin) AddTable() {}
