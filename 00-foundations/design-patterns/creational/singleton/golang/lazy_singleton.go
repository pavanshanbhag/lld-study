package singleton

import "sync"

// LazySingleton implements lazy initialization singleton pattern.
type LazySingleton struct{}

var (
	lazyInstance *LazySingleton
	lazyOnce     sync.Once
)

func GetLazyInstance() *LazySingleton {
	lazyOnce.Do(func() {
		lazyInstance = &LazySingleton{}
	})
	return lazyInstance
}
