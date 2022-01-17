package main

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		// num of keys to track frequency, usually 10*MaxCost
		NumCounters: 100,
		// cache size(max num of items)
		MaxCost: 10,
		// number of keys per Get buffer
		BufferItems: 64,
		// !important: always set true if not limiting memory
		IgnoreInternalCost: true,
	})
	if err != nil {
		panic(err)
	}

	// put 20(>10) items to cache
	for i := 0; i < 20; i++ {
		cache.Set(i, i, 1)
	}

	// wait for value to pass through buffers
	cache.Wait()

	cntCacheMiss := 0
	for i := 0; i < 20; i++ {
		if _, ok := cache.Get(i); !ok {
			cntCacheMiss++
		}
	}
	fmt.Printf("%d of 20 items missed\n", cntCacheMiss)
}
