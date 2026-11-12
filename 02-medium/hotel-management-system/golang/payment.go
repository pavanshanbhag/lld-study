package hotelmanagement

import "fmt"

type Payment interface {
	ProcessPayment(amount float64) error
}

type CreditCardPayment struct{}

func NewCreditCardPayment() *CreditCardPayment {
	return &CreditCardPayment{}
}

func (p *CreditCardPayment) ProcessPayment(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid payment amount: %v", amount)
	}
	return nil
}

type CashPayment struct{}

func NewCashPayment() *CashPayment {
	return &CashPayment{}
}

func (p *CashPayment) ProcessPayment(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid payment amount: %v", amount)
	}
	return nil
}
