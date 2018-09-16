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
	m.mutex.RLock()
	item, ok := m.items[key]
	if !ok {
		err = gocache.ErrNotMatch
	} else {
		if !item.isExpired() {
			ret = item.data
		}
	}
	m.mutex.RUnlock()
	return
}

func (m *MemCache) Set(key string, val interface{}, opts ...time.Time) error {
	ex := time.Now().Add(m.ttl)
	item := Item{data: &val, expires: &ex}
	m.mutex.Lock()
	m.items[key] = &item
	m.mutex.Unlock()
	return nil
}

func (m *MemCache) Remove(key string) error {
	m.mutex.Lock()
	_, ok := m.items[key]
	if ok {
		delete(m.items, key)
	}
	m.mutex.Unlock()
	return nil
}

func (m *MemCache) IsExpired(key string) (ret bool) {
	m.mutex.RLock()
	item, ok := m.items[key]
	ret = true
	if ok {
		ret = item.isExpired()
		if ret == true {
			delete(m.items, key)
		}
	}
	m.mutex.RUnlock()
	return
}

func (m MemCache) String() string {
	m.mutex.RLock()
	ts := ""
	for k, v := range m.items {
		ts += "[" + k + "]" + (*v).String() + ","
	}
	m.mutex.RUnlock()
	return fmt.Sprintf("MemCache{ttl:%d,items:%v}", m.ttl, ts)
}
