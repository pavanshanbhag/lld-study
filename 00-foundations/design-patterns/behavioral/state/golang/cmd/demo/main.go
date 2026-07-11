package main

import "state"

func main() {
	vm := state.NewVendingMachine()

	vm.InsertCoin(1.0)

	vm.SelectItem("Soda")

	vm.InsertCoin(1.5)

	vm.DispenseItem()

	vm.SelectItem("Chips")
	vm.InsertCoin(2.0)
	vm.DispenseItem()
}
