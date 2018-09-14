package memory

import (
	"testing"
	"time"
)

func now() (ret *time.Time) {
	rett := time.Now()
	ret = &rett
	return
}

func TestItem_IsExpired(t *testing.T) {
	item := Item{expires: now()}
	isexpired := item.isExpired()
	if !isexpired {
		t.Error("Err should expired", item)
	}
	item.touch(2000)
	isexpired = item.isExpired()
	if isexpired {
		t.Error("Err should not expired", item)
	}

	t1 := time.Now().Add(10000)
	item2 := Item{expires: &t1}
	if item2.isExpired() {
		t.Error(item2)
	}
}
