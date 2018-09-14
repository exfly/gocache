package memory

import (
	"testing"

	"github.com/ExFly/gocache"
)

func TestMemCache_Cache(t *testing.T) {
	cacher := NewMemCache()
	cacher.Set("key1", "val1")
	cacher.Set("key2", "val2")
	t.Log(cacher)
	cacher.Remove("key2")
	data, err := cacher.Get("key2")
	t.Log(cacher)
	t.Logf("get key had del: %v", data)
	if err != gocache.ErrNotMatch {
		t.Error(err)
	}
}
