package admin

import (
	"encoding/json"
	"errors"

	"github.com/Alienero/Rambo/meta"
)

var (
	errInvalidUser = errors.New("invalid user")
	errDBExisted   = errors.New("db is arlready existed")
	errMarshlJSON  = errors.New("marshl json error")
)

// Admin provide some method to manage database
type Admin struct {
	isLogin bool
}

// AddUser add a database user
func (Admin) AddUser(user, password string) error {
	return meta.Meta.AddUser(user, password)
}

// GetUser will get user's info.
func (Admin) GetUser(user string) (string, error) {
	resp, err := meta.Meta.GetUserInfo(user)
	if err != nil {
		return "", err
	}
	data, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		return "", errMarshlJSON
	}
	return string(data), nil
}

// Login check the user.
func (m *Admin) Login(user, password string) bool {
	m.isLogin = meta.Meta.CheckUserDirect(user, password)
	return m.isLogin
}
