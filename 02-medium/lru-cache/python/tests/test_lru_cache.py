from lru_cache import LRUCache


def test_get_put_eviction() -> None:
    cache: LRUCache[int, str] = LRUCache(2)
    cache.put(1, "one")
    cache.put(2, "two")
    assert cache.get(1) == "one"

    cache.put(3, "three")
    assert cache.get(2) is None
    assert cache.get(3) == "three"


def test_update_existing_key() -> None:
    cache: LRUCache[int, str] = LRUCache(2)
    cache.put(1, "one")
    cache.put(2, "two")
    cache.put(1, "ONE")

    assert cache.get(1) == "ONE"
    assert cache.size() == 2


def test_remove_and_clear() -> None:
    cache: LRUCache[int, str] = LRUCache(3)
    cache.put(1, "one")
    cache.put(2, "two")
    cache.remove(1)
    assert cache.get(1) is None
    assert cache.size() == 1

    cache.clear()
    assert cache.size() == 0
    assert cache.get(2) is None
