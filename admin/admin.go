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
	info    *meta.Info
}

// NewAdmin get a new Admin instance
func NewAdmin(info *meta.Info) *Admin {
	return &Admin{
		info: info,
	}
}

// AddUser add a database user
func (a *Admin) AddUser(user, password string) error {
	return a.info.AddUser(user, password)
}

// GetUser will get user's info.
func (a *Admin) GetUser(user string) (string, error) {
	resp, err := a.info.GetUserInfo(user)
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
func (a *Admin) Login(user, password string) bool {
	a.isLogin = a.info.CheckUserDirect(user, password)
	return a.isLogin
}
