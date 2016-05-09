package server

import (
	"sync"

	"github.com/Alienero/Rambo/mysql/client"
)

type Bpool struct {
	bakens map[string]*client.DB
	sync.RWMutex
	isOpen bool
}

func (p *Bpool) GetConnect(user, db string, num int) {
	// key := fmt.Sprintf("%s_%s_%d", user, db, num)
	// if p.isOpen {

	// }
}
