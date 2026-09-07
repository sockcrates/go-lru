package lru

type LRU[K comparable, V any] struct {
	_ struct{}
}

func New[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{}
}

func (l *LRU[K, V]) Put(key K, value V) {}

func (l *LRU[K, V]) Get(key K) (V, bool) {
	var zero V
	return zero, false
}
