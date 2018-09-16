package memory

import (
	"testing"

	"github.com/ExFly/gocache"
)

func TestMemCache_Cache(t *testing.T) {
	cacher := NewMemCache(2, 1000000000)
	cacher.Set("key1", "val1")
	cacher.Set("key2", "val2")
	t.Log(cacher)
	data, _ := cacher.Get("key2")
	t.Log(data)
	if val, ok := data.(string); !ok || val != "val2" {
		t.Error("val2!=", data, ok)
	}
	cacher.Remove("key2")
	data, err := cacher.Get("key2")
	t.Log(cacher)
	t.Logf("get key had del: %v", data)
	if err != gocache.ErrNotMatch {
		t.Error(err)
	}
}
