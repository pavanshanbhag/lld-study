package singleton

// EagerSingleton implements eager initialization singleton pattern.
type EagerSingleton struct{}

var eagerInstance = &EagerSingleton{}

func GetEagerInstance() *EagerSingleton {
	return eagerInstance
}
