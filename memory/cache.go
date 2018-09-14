package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/ExFly/gocache"
)

type MemCache struct {
	mutex sync.RWMutex
	ttl   time.Duration
	items map[string]*Item
}

func NewMemCache() gocache.Cache {
	return &MemCache{
		ttl:   1000,
		items: make(map[string]*Item),
	}
}

func (m *MemCache) Get(key string) (ret interface{}, err error) {
	item, ok := m.items[key]
	if !ok {
		err = gocache.ErrNotMatch
		return
	} else {
		if !item.isExpired() {
			ret = item.data
		}
	}
	return
}

func (m *MemCache) Set(key string, val interface{}, opts ...time.Time) error {
	ex := time.Now().Add(m.ttl)
	item := Item{data: &val, expires: &ex}
	m.items[key] = &item
	return nil
}
func (m *MemCache) Remove(key string) error {
	_, ok := m.items[key]
	if ok {
		delete(m.items, key)
	}
	return nil
}
func (m *MemCache) IsExpired(key string) (ret bool) {
	item, ok := m.items[key]
	if !ok {
		ret = true
	} else {
		ret = item.isExpired()
		delete(m.items, key)
	}
	return
}
func (m MemCache) String() string {
	ts := ""
	for k, v := range m.items {
		ts += "[" + k + "]" + (*v).String() + ","
	}
	return fmt.Sprintf("MemCache{ttl:%d,items:%v}", m.ttl, ts)
}
