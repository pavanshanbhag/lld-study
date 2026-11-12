package digitalwallet

import "math/big"

type PaymentMethod interface {
	ProcessPayment(amount *big.Float, currency Currency) error
	ID() string
	User() *User
}

type BasePaymentMethod struct {
	id   string
	user *User
}

func NewBasePaymentMethod(id string, user *User) BasePaymentMethod {
	return BasePaymentMethod{id: id, user: user}
}

func (b *BasePaymentMethod) ID() string {
	return b.id
}

func (b *BasePaymentMethod) User() *User {
	return b.user
}
