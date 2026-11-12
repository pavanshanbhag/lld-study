package vending_machine

import "testing"

func TestNewVendingMachine(t *testing.T) {
	t.Parallel()
	vm := NewVendingMachine()
	if vm == nil || vm.currentState == nil {
		t.Fatal("expected initialized vending machine")
	}
}

func TestVendingMachinePurchaseFlow(t *testing.T) {
	t.Parallel()

	vm := NewVendingMachine()
	coke := NewProduct("Coke", 1.0)
	vm.inventory.AddProduct(coke, 2)

	if err := vm.SelectProduct(coke); err != nil {
		t.Fatalf("select: %v", err)
	}
	if err := vm.InsertCoin(QUARTER); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := vm.InsertCoin(QUARTER); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := vm.InsertCoin(QUARTER); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := vm.InsertCoin(QUARTER); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := vm.DispenseProduct(); err != nil {
		t.Fatalf("dispense: %v", err)
	}
	if err := vm.ReturnChange(); err != nil {
		t.Fatalf("return change: %v", err)
	}
	if got := vm.inventory.Quantity(coke); got != 1 {
		t.Fatalf("quantity = %d, want 1", got)
	}
}

func TestVendingMachineUnavailableProduct(t *testing.T) {
	t.Parallel()

	vm := NewVendingMachine()
	water := NewProduct("Water", 1.0)
	if err := vm.SelectProduct(water); err == nil {
		t.Fatal("expected error for unavailable product")
	}
}
