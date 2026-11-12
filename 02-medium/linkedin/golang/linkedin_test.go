package linkedin

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewProfile()
}
