package main

import (
	"fmt"

	"singleton"
)

func main() {
	lazyA := singleton.GetLazyInstance()
	lazyB := singleton.GetLazyInstance()
	fmt.Printf("lazy singleton same instance: %v\n", lazyA == lazyB)

	eager := singleton.GetEagerInstance()
	fmt.Printf("eager singleton: %p\n", eager)

	ts := singleton.GetThreadSafeInstance()
	fmt.Printf("thread-safe singleton: %p\n", ts)

	dc := singleton.GetDoubleCheckedInstance()
	fmt.Printf("double-checked singleton: %p\n", dc)
}
