package coffeevendingmachine

import "testing"

func TestNewCoffeeMachine(t *testing.T) {
	t.Parallel()
	cm := NewCoffeeMachine()
	if cm == nil || len(cm.ingredients) == 0 {
		t.Fatal("expected initialized coffee machine")
	}
}
