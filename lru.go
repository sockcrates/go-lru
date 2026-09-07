package lru

type LRU[K comparable, V any] struct {
	_ struct {}
}

func New[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{}
}
