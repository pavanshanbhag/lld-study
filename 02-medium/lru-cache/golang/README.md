# LRU Cache — Go (modernized)

Idiomatic Go implementation:

- Generic `NewLRUCache[K, V](capacity)` (no singleton)
- O(1) get/put via map + doubly linked list
- `sync.RWMutex` for concurrent access
- `log/slog` in demo
- Table-driven tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./02-medium/lru-cache/golang/...
```
