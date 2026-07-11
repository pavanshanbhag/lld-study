package facade

import "time"

func simulateDelay() {
	time.Sleep(100 * time.Millisecond)
}
