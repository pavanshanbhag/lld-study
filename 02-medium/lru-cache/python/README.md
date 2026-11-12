# LRU Cache — Python (modernized)

Thread-safe LRU cache using a hash map + doubly linked list:

- Generic `LRUCache[K, V]`
- O(1) get/put/remove
- Constructor-based (no singleton)
- Demo uses stdout for UX

## Run

```bash
# from repo root
uv run python 02-medium/lru-cache/python/lru_cache_demo.py
uv run pytest 02-medium/lru-cache/python/tests -q
```
