package vending_machine

import "fmt"

type VendingMachineState interface {
	SelectProduct(product *Product) error
	InsertCoin(coin Coin) error
	InsertNote(note Note) error
	DispenseProduct() error
	ReturnChange() error
}

type IdleState struct {
	vendingMachine *VendingMachine
}

func (s *IdleState) SelectProduct(product *Product) error {
	if !s.vendingMachine.inventory.IsAvailable(product) {
		return fmt.Errorf("product not available: %s", product.name)
	}
	s.vendingMachine.selectedProduct = product
	s.vendingMachine.SetState(s.vendingMachine.readyState)
	return nil
}

func (s *IdleState) InsertCoin(coin Coin) error {
	return fmt.Errorf("please select a product first")
}
func (s *IdleState) InsertNote(note Note) error {
	return fmt.Errorf("please select a product first")
}
func (s *IdleState) DispenseProduct() error {
	return fmt.Errorf("please select a product and make payment")
}
func (s *IdleState) ReturnChange() error {
	return fmt.Errorf("no change to return")
}

type ReadyState struct {
	vendingMachine *VendingMachine
}

func (s *ReadyState) SelectProduct(product *Product) error {
	return fmt.Errorf("product already selected; please make payment")
}

func (s *ReadyState) InsertCoin(coin Coin) error {
	s.vendingMachine.totalPayment += float64(coin)
	s.checkPaymentStatus()
	return nil
}

func (s *ReadyState) InsertNote(note Note) error {
	s.vendingMachine.totalPayment += float64(note)
	s.checkPaymentStatus()
	return nil
}

func (s *ReadyState) DispenseProduct() error {
	return fmt.Errorf("please make payment first")
}

func (s *ReadyState) ReturnChange() error {
	change := s.vendingMachine.totalPayment - s.vendingMachine.selectedProduct.price
	s.vendingMachine.ResetPayment()
	s.vendingMachine.SetState(s.vendingMachine.idleState)
	if change <= 0 {
		return nil
	}
	return nil
}

func (s *ReadyState) checkPaymentStatus() {
	if s.vendingMachine.totalPayment >= s.vendingMachine.selectedProduct.price {
		s.vendingMachine.SetState(s.vendingMachine.dispenseState)
	}
}

type DispenseState struct {
	vendingMachine *VendingMachine
}

func (s *DispenseState) SelectProduct(product *Product) error {
	return fmt.Errorf("product already selected")
}
func (s *DispenseState) InsertCoin(coin Coin) error {
	return fmt.Errorf("please collect the product")
}
func (s *DispenseState) InsertNote(note Note) error {
	return fmt.Errorf("please collect the product")
}
func (s *DispenseState) DispenseProduct() error {
	product := s.vendingMachine.selectedProduct
	if err := s.vendingMachine.inventory.Dispense(product); err != nil {
		return err
	}
	s.vendingMachine.SetState(s.vendingMachine.returnChangeState)
	return nil
}
func (s *DispenseState) ReturnChange() error {
	return fmt.Errorf("please collect the product first")
}

type ReturnChangeState struct {
	vendingMachine *VendingMachine
}

func (s *ReturnChangeState) SelectProduct(product *Product) error {
	return fmt.Errorf("please collect the change first")
}
func (s *ReturnChangeState) InsertCoin(coin Coin) error {
	return fmt.Errorf("please collect the change first")
}
func (s *ReturnChangeState) InsertNote(note Note) error {
	return fmt.Errorf("please collect the change first")
}
func (s *ReturnChangeState) DispenseProduct() error {
	return fmt.Errorf("product already dispensed; please collect the change")
}
func (s *ReturnChangeState) ReturnChange() error {
	change := s.vendingMachine.totalPayment - s.vendingMachine.selectedProduct.price
	s.vendingMachine.ResetPayment()
	s.vendingMachine.ResetSelectedProduct()
	s.vendingMachine.SetState(s.vendingMachine.idleState)
	_ = change
	return nil
}
