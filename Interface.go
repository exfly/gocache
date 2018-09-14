package gocache

import (
	"errors"
	"time"
)

var (
	ErrNotMatch = errors.New("Cache not match")
)

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, val interface{}, opts ...time.Time) error
	Remove(key string) error
	IsExpired(key string) bool
}
