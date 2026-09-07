package lru_test

import (
	"fmt"

	"github.com/sockcrates/go-lru"
)

func Empty() {
	lru.New[string, int](2)
	fmt.Println("(empty lru)")
}
