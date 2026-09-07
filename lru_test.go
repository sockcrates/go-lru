package lru_test

import (
	"fmt"
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
	if ok == true {
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
