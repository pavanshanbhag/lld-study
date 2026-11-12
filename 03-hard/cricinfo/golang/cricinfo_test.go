package cricinfo

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewCricinfoSystem()
}
