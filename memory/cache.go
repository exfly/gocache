package memory

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ExFly/gocache"
	"github.com/ExFly/gocache/memory/simplelru"
)

type MemCache struct {
	mutex sync.RWMutex
	ttl   time.Duration
	items simplelru.LRUCache
}

func NewMemCache(size int, ttl time.Duration) gocache.Cache {
	l, err := simplelru.NewLRU(size, nil)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return &MemCache{
		ttl:   ttl,
		items: l,
	}
}

func (m *MemCache) Get(key string) (ret interface{}, err error) {
	m.mutex.RLock()
	item, ok := m.items.Get(key)
	if !ok {
		err = gocache.ErrNotMatch
	} else {
		it, _ := item.(*Item)
		if !it.isExpired() {
			datap := it.data
			if datap != nil {
				ret = *datap
			}
		} else {
			log.Printf("key:%v is expired", key)
		}
	}
	m.mutex.RUnlock()
	return
}

func (m *MemCache) Set(key string, val interface{}, opts ...time.Time) error {
	expired := time.Now().Add(m.ttl)
	item := Item{data: &val, expires: &expired}
	m.mutex.Lock()
	m.items.Add(key, &item)
	m.mutex.Unlock()
	return nil
}

func (m *MemCache) Remove(key string) error {
	m.mutex.Lock()
	ok := m.items.Remove(key)
	if !ok {
		return gocache.ErrNotFoundItem
	}
	m.mutex.Unlock()
	return nil
}

func (m *MemCache) IsExpired(key string) (ret bool) {
	m.mutex.RLock()
	item, ok := m.items.Get(key)
	ret = true
	if ok {
		it := item.(*Item)
		ret = it.isExpired()
		if ret == true {
			m.items.Remove(key)
		}
	}
	m.mutex.RUnlock()
	return
}

func (m MemCache) String() string {
	m.mutex.RLock()
	ts := ""
	for _, k := range m.items.Keys() {
		it, ok := m.items.Get(k)
		if !ok {
			continue
		}
		ts += "[" + k.(string) + "]" + it.(*Item).String() + ","
	}
	m.mutex.RUnlock()
	return fmt.Sprintf("MemCache{ttl:%d,items:%v}", m.ttl, ts)
}
