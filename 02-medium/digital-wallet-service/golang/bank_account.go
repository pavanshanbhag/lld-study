package digitalwallet

import "math/big"

type BankAccount struct {
	BasePaymentMethod
	AccountNumber string
	RoutingNumber string
}

func NewBankAccount(id string, user *User, accountNumber, routingNumber string) *BankAccount {
	return &BankAccount{
		BasePaymentMethod: NewBasePaymentMethod(id, user),
		AccountNumber:     accountNumber,
		RoutingNumber:     routingNumber,
	}
}

func (b *BankAccount) ProcessPayment(amount *big.Float, currency Currency) error {
	return nil
}
