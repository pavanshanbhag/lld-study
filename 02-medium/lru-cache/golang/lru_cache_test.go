package lrucache

import (
	"fmt"
	"sync"
	"testing"
)

func TestLRUCacheGetPutEviction(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(2, "two")

	if val, ok := cache.Get(1); !ok || val != "one" {
		t.Fatalf("Get(1) = (%q, %v), want (one, true)", val, ok)
	}

	cache.Put(3, "three") // evicts key 2 (least recently used)

	if _, ok := cache.Get(2); ok {
		t.Fatal("Get(2) should miss after eviction")
	}
	if val, ok := cache.Get(3); !ok || val != "three" {
		t.Fatalf("Get(3) = (%q, %v), want (three, true)", val, ok)
	}
}

func TestLRUCacheUpdateExisting(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Put(1, "ONE")

	if val, ok := cache.Get(1); !ok || val != "ONE" {
		t.Fatalf("Get(1) = (%q, %v), want (ONE, true)", val, ok)
	}
	if cache.Size() != 2 {
		t.Fatalf("Size() = %d, want 2", cache.Size())
	}
}

func TestLRUCacheClear(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[int, string](3)
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Clear()

	if cache.Size() != 0 {
		t.Fatalf("Size() after Clear = %d, want 0", cache.Size())
	}
	if _, ok := cache.Get(1); ok {
		t.Fatal("Get(1) should miss after Clear")
	}
}

func TestLRUCacheConcurrentAccess(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[int, int](100)
	var wg sync.WaitGroup

	for i := range 50 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cache.Put(n, n*10)
		}(i)
	}

	for i := range 50 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			_, _ = cache.Get(n)
		}(i)
	}

	wg.Wait()

	if cache.Size() > 100 {
		t.Fatalf("Size() = %d, exceeds capacity 100", cache.Size())
	}
}

func TestLRUCacheMissZeroValue(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[string, int](1)
	val, ok := cache.Get("missing")
	if ok {
		t.Fatalf("Get(missing) = (%d, true), want miss", val)
	}
	if val != 0 {
		t.Fatalf("zero value = %d, want 0", val)
	}
}

func TestLRUCacheCapacityOne(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[int, string](1)
	cache.Put(1, "a")
	cache.Put(2, "b")

	if _, ok := cache.Get(1); ok {
		t.Fatal("key 1 should be evicted")
	}
	if val, ok := cache.Get(2); !ok || val != "b" {
		t.Fatalf("Get(2) = (%q, %v), want (b, true)", val, ok)
	}
}

func ExampleNewLRUCache() {
	cache := NewLRUCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Get(1)
	cache.Put(3, "three")

	if val, ok := cache.Get(2); ok {
		fmt.Println(val)
	} else {
		fmt.Println("evicted")
	}
	// Output: evicted
}
