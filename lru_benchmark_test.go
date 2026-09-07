package lru_test

import (
	"testing"

	"github.com/sockcrates/go-lru"
)

const benchmarkCapacity = 1024

func newFullBenchmarkCache() *lru.LRU[int, int] {
	c := lru.New[int, int](benchmarkCapacity)
	for key := range benchmarkCapacity {
		c.Put(key, key)
	}
	return c
}

func BenchmarkLRUGetHotHit(b *testing.B) {
	c := newFullBenchmarkCache()
	b.ReportAllocs()

	for b.Loop() {
		c.Get(0)
	}
}

func BenchmarkLRUGetRotatingHit(b *testing.B) {
	c := newFullBenchmarkCache()
	key := 0
	b.ReportAllocs()

	for b.Loop() {
		c.Get(key)
		key = (key + 1) & (benchmarkCapacity - 1)
	}
}

func BenchmarkLRUGetMiss(b *testing.B) {
	c := newFullBenchmarkCache()
	b.ReportAllocs()

	for b.Loop() {
		c.Get(-1)
	}
}

func BenchmarkLRUPutUpdate(b *testing.B) {
	c := newFullBenchmarkCache()
	value := 0
	b.ReportAllocs()

	for b.Loop() {
		c.Put(0, value)
		value++
	}
}

func BenchmarkLRUPutEvict(b *testing.B) {
	c := newFullBenchmarkCache()
	key := benchmarkCapacity
	b.ReportAllocs()

	for b.Loop() {
		c.Put(key, key)
		key++
	}
}

func BenchmarkLRUConcurrentGetHotHit(b *testing.B) {
	c := newFullBenchmarkCache()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get(0)
		}
	})
}
