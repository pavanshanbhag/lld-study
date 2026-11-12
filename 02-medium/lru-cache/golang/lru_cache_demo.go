package lrucache

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	lruCache := NewLRUCache[int, string](3)

	lruCache.Put(1, "Value 1")
	lruCache.Put(2, "Value 2")
	lruCache.Put(3, "Value 3")

	if val, exists := lruCache.Get(1); exists {
		logger.Info("cache hit", "key", 1, "value", val)
	}
	if val, exists := lruCache.Get(2); exists {
		logger.Info("cache hit", "key", 2, "value", val)
	}

	lruCache.Put(4, "Value 4")

	if _, exists := lruCache.Get(3); !exists {
		logger.Info("key evicted", "key", 3)
	}

	if val, exists := lruCache.Get(4); exists {
		logger.Info("cache hit", "key", 4, "value", val)
	}

	lruCache.Put(2, "Updated Value 2")

	if val, exists := lruCache.Get(1); exists {
		logger.Info("cache hit", "key", 1, "value", val)
	}
	if val, exists := lruCache.Get(2); exists {
		logger.Info("cache hit", "key", 2, "value", val)
	}
}
