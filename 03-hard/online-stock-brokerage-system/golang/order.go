package onlinestockbrokeragesystem

type Order interface {
	Execute() error
	Status() OrderStatus
	SetStatus(status OrderStatus)
}

type BaseOrder struct {
	OrderID  string
	Account  *Account
	Stock    *Stock
	Quantity int
	Price    float64
	status   OrderStatus
}

func (o *BaseOrder) Status() OrderStatus {
	return o.status
}

func (o *BaseOrder) SetStatus(status OrderStatus) {
	o.status = status
}
