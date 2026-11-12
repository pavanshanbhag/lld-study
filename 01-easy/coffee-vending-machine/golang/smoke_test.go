package coffeevendingmachine

import "testing"

func TestNewCoffeeMachine(t *testing.T) {
	t.Parallel()
	cm := NewCoffeeMachine()
	if cm == nil || len(cm.ingredients) == 0 {
		t.Fatal("expected initialized coffee machine")
	}
}

func TestDispenseEspresso(t *testing.T) {
	t.Parallel()
	cm := NewCoffeeMachine()
	coffee, err := cm.SelectCoffee("Espresso")
	if err != nil {
		t.Fatalf("select: %v", err)
	}
	change, err := cm.DispenseCoffee(coffee, NewPayment(3.0))
	if err != nil {
		t.Fatalf("dispense: %v", err)
	}
	if change != 0.5 {
		t.Fatalf("change = %v, want 0.5", change)
	}
	if cm.IngredientQuantity("Coffee") != 9 {
		t.Fatalf("coffee qty = %d, want 9", cm.IngredientQuantity("Coffee"))
	}
}

func TestSelectUnknownCoffee(t *testing.T) {
	t.Parallel()
	cm := NewCoffeeMachine()
	if _, err := cm.SelectCoffee("Mocha"); err == nil {
		t.Fatal("expected error for unknown coffee")
	}
}
