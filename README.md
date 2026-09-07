# go-lru

`go-lru` is a small generic, concurrency-safe least-recently-used cache for
Go. It supports any comparable key type and any value type, and evicts the
least recently accessed entry when the cache exceeds its capacity.

## Usage

Install the module:

```console
$ go get github.com/sockcrates/go-lru
```

Create a cache with a fixed capacity, add values with `Put`, and retrieve them
with `Get`:

```go
package main

import (
	"fmt"

	"github.com/sockcrates/go-lru"
)

func main() {
	cache := lru.New[string, int](2)
	cache.Put("answer", 42)

	if value, ok := cache.Get("answer"); ok {
		fmt.Println(value)
	}
}
```

`Get` returns the value and `true` for a hit. A miss returns the zero value of
the cache's value type and `false`, so the boolean must be used to distinguish
a missing key from a stored zero value.

## Requirements and development

The module targets the Go version declared in [`go.mod`](go.mod). The
[mise](https://mise.jdx.dev/) configuration installs the required Go toolchain.

```console
$ mise install
$ go test ./...              # Run the test suite
$ go test -race ./...        # Check concurrent access with the race detector
$ go test -bench=. ./...     # Run the benchmarks
$ gofmt -d .                 # Check formatting
$ go vet ./...               # Run static analysis
```

## API documentation

The public API can be viewed locally with `go doc`:

```console
$ go doc github.com/sockcrates/go-lru
$ go doc github.com/sockcrates/go-lru.LRU
```

## Behaviour and performance

Both successful `Get` calls and `Put` calls mark an entry as most recently
used. Updating an existing key replaces its value without increasing the cache
size. Adding a new key to a full cache evicts the least recently used entry. A
cache with zero or negative capacity retains no entries.

`Get` and `Put` are constant time, and all access is serialized by a mutex, so
a cache may be shared safely between goroutines. Keys and values are stored
directly; the cache does not copy or manage resources referenced by them.
