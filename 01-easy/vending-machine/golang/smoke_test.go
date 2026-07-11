package vending_machine

import "testing"

func TestNewVendingMachine(t *testing.T) {
	t.Parallel()
	vm := NewVendingMachine()
	if vm == nil || vm.currentState == nil {
		t.Fatal("expected initialized vending machine")
	}
}
