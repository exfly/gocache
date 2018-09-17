package memory

import (
	"testing"
	"time"

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
func Benchmark_Set(b *testing.B) {
	cacher := NewMemCache(2, 5*time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cacher.Set("key"+string(i), "val"+string(i))
	}
}

func Benchmark_SetParallel(b *testing.B) {
	cacher := NewMemCache(2, 5*time.Second)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cacher.Set("key1", "val1")
		}
	})
}

func Benchmark_Get(b *testing.B) {
	cacher := NewMemCache(2, 5*time.Second)
	cacher.Set("key1", "val1")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cacher.Get("key1")
	}
}

// 测试并发效率
func Benchmark_GetParallel(b *testing.B) {
	cacher := NewMemCache(2, 5*time.Second)
	cacher.Set("key1", "val1")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cacher.Get("key1")
		}
	})
}
