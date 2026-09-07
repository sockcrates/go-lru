package lru_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/sockcrates/go-lru"
)

func TestEmpty(t *testing.T) {
	lru.New[string, int](2)
	fmt.Println("(empty lru)")
}

func TestLRUGetsEmpty(t *testing.T) {
	c := lru.New[string, int](2)
	v, ok := c.Get("miss")
	if ok == true || v != 0 {
		t.Errorf("expected no value and false, got %v and %v", v, ok)
	}
}

func TestLRUPutsAndGets(t *testing.T) {
	c := lru.New[string, int](2)
	c.Put("value-1", 1)
	v, ok := c.Get("value-1")
	if ok != true || v != 1 {
		t.Errorf("expected value of 1 and true, got %v and %v", v, ok)
	}
}

func TestLRUStoresZeroValue(t *testing.T) {
	c := lru.New[string, int](2)
	c.Put("value-1", 0)
	v, ok := c.Get("value-1")
	if ok != true || v != 0 {
		t.Errorf("expected value of 0 and true, got %v and %v", v, ok)
	}
}

func TestLRUUpdatesExistingValue(t *testing.T) {
	c := lru.New[string, int](2)
	c.Put("value-1", 1)
	c.Put("value-1", 2)
	v, ok := c.Get("value-1")
	if ok != true || v != 2 {
		t.Errorf("expected value of 2 and true, got %v and %v", v, ok)
	}
}

func TestLRUEvictsLeastRecentlyUsed(t *testing.T) {
	c := lru.New[string, int](2)
	c.Put("value-1", 1)
	c.Put("value-2", 2)
	c.Put("value-3", 3)
	v, ok := c.Get("value-1")
	if ok == true {
		t.Errorf("expected no value and false, got %v and %v", v, ok)
	}
	v, ok = c.Get("value-3")
	if ok != true || v != 3 {
		t.Errorf("expected value of 3 and true, got %v and %v", v, ok)
	}
}

func TestLRUGetUpdatesRecency(t *testing.T) {
	c := lru.New[string, int](2)
	c.Put("value-1", 1)
	c.Put("value-2", 2)
	c.Get("value-1")
	c.Put("value-3", 3)
	v, ok := c.Get("value-2")
	if ok == true {
		t.Errorf("expected no value and false, got %v and %v", v, ok)
	}
	v, ok = c.Get("value-1")
	if ok != true || v != 1 {
		t.Errorf("expected value of 1 and true, got %v and %v", v, ok)
	}
}

func TestLRUPutUpdatesRecency(t *testing.T) {
	c := lru.New[string, int](2)
	c.Put("value-1", 1)
	c.Put("value-2", 2)
	c.Put("value-1", 3)
	c.Put("value-3", 4)
	v, ok := c.Get("value-2")
	if ok == true {
		t.Errorf("expected no value and false, got %v and %v", v, ok)
	}
	v, ok = c.Get("value-1")
	if ok != true || v != 3 {
		t.Errorf("expected value of 3 and true, got %v and %v", v, ok)
	}
}

func TestLRUWithZeroCapacity(t *testing.T) {
	c := lru.New[string, int](0)
	c.Put("value-1", 1)
	v, ok := c.Get("value-1")
	if ok == true || v != 0 {
		t.Errorf("expected no value and false, got %v and %v", v, ok)
	}
}

func TestLRUConcurrentPutsAndGets(t *testing.T) {
	c := lru.New[int, int](16)
	start := make(chan struct{})
	var wg sync.WaitGroup

	for worker := 0; worker < 32; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for value := 0; value < 100; value++ {
				c.Put(value, value)
				c.Get(value)
			}
		}()
	}

	close(start)
	wg.Wait()

	c.Put(-1, -1)
	v, ok := c.Get(-1)
	if ok != true || v != -1 {
		t.Errorf("expected value of -1 and true, got %v and %v", v, ok)
	}
}
