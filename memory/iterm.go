package memory

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	sync.RWMutex
	data    *interface{}
	expires *time.Time
}

func (it *Item) touch(duration time.Duration) {
	expiration := time.Now().Add(duration)
	it.Lock()
	it.expires = &expiration
	it.Unlock()
}

func (it *Item) isExpired() (ret bool) {
	if it.expires != nil {
		it.RLock()
		ret = it.expires.Before(time.Now())
		it.RUnlock()
	}
	return
}
func (itt Item) String() (ret string) {
	return fmt.Sprintf("Item{data:%v,expires:%v}", *itt.data, itt.expires)
}
