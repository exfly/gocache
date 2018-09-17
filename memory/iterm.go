package memory

import (
	"fmt"
	"time"
)

// Item concurrent unsafe
type Item struct {
	data    *interface{}
	expires *time.Time
}

func (it *Item) touch(duration time.Duration) {
	expiration := time.Now().Add(duration)
	it.expires = &expiration
}

func (it *Item) isExpired() (ret bool) {
	if it.expires != nil {
		ret = it.expires.Before(time.Now())
	}
	return
}
func (itt Item) String() (ret string) {
	return fmt.Sprintf("Item{data:%v,expires:%v}", *itt.data, itt.expires)
}
