package onlineshopping

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewShoppingCart()
}
