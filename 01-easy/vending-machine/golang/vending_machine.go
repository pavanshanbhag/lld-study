package vending_machine

type VendingMachine struct {
	inventory         *Inventory
	idleState         VendingMachineState
	readyState        VendingMachineState
	dispenseState     VendingMachineState
	returnChangeState VendingMachineState
	currentState      VendingMachineState
	selectedProduct   *Product
	totalPayment      float64
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		inventory: NewInventory(),
	}
	vm.idleState = &IdleState{vm}
	vm.readyState = &ReadyState{vm}
	vm.dispenseState = &DispenseState{vm}
	vm.returnChangeState = &ReturnChangeState{vm}
	vm.currentState = vm.idleState
	return vm
}

func (vm *VendingMachine) SelectProduct(product *Product) error {
	return vm.currentState.SelectProduct(product)
}

func (vm *VendingMachine) InsertCoin(coin Coin) error {
	return vm.currentState.InsertCoin(coin)
}

func (vm *VendingMachine) InsertNote(note Note) error {
	return vm.currentState.InsertNote(note)
}

func (vm *VendingMachine) DispenseProduct() error {
	return vm.currentState.DispenseProduct()
}

func (vm *VendingMachine) ReturnChange() error {
	return vm.currentState.ReturnChange()
}

func (vm *VendingMachine) SetState(state VendingMachineState) {
	vm.currentState = state
}

func (vm *VendingMachine) ResetPayment() {
	vm.totalPayment = 0
}

func (vm *VendingMachine) ResetSelectedProduct() {
	vm.selectedProduct = nil
}

func (vm *VendingMachine) TotalPayment() float64 {
	return vm.totalPayment
}

func (vm *VendingMachine) SelectedProduct() *Product {
	return vm.selectedProduct
}
