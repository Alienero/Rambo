package uuid

import "github.com/twinj/uuid"

// Get will get a uuid
func Get() string {
	return uuid.NewV4().String()
}
