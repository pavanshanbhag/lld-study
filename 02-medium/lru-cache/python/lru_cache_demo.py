"""LRU cache demo entrypoint."""

from __future__ import annotations

import logging

from lru_cache import LRUCache

logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
logger = logging.getLogger(__name__)


def main() -> None:
    cache: LRUCache[str, int] = LRUCache(3)

    cache.put("a", 1)
    cache.put("b", 2)
    cache.put("c", 3)

    logger.info("get(a)=%s", cache.get("a"))

    cache.put("d", 4)
    logger.info("get(b)=%s (evicted)", cache.get("b"))


if __name__ == "__main__":
    main()
