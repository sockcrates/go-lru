package lru

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRU[K comparable, V any] struct {
	cache    map[K]*list.Element
	capacity int
	list     *list.List
	mu       sync.Mutex
	_        struct{}
}

func New[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{cache: make(map[K]*list.Element), capacity: capacity, list: list.New()}
}

func (l *LRU[K, V]) Put(key K, value V) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.cache[key]; ok {
		e.Value = entry[K, V]{key: key, value: value}
		l.list.MoveToFront(e)
		return
	}
	e := l.list.PushFront(entry[K, V]{key: key, value: value})
	l.cache[key] = e
	if l.list.Len() > l.capacity {
		lru := l.list.Back()
		data := lru.Value
		delete(l.cache, data.(entry[K, V]).key)
		l.list.Remove(lru)
	}
}

func (l *LRU[K, V]) Get(key K) (V, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	d, ok := l.cache[key]
	if !ok {
		var zero V
		return zero, false
	}
	l.list.MoveToFront(d)
	return d.Value.(entry[K, V]).value, true
}
