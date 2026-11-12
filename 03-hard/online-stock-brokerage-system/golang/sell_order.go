package onlinestockbrokeragesystem

type SellOrder struct {
	BaseOrder
}

func NewSellOrder(orderID string, account *Account, stock *Stock, quantity int, price float64) *SellOrder {
	return &SellOrder{
		BaseOrder: BaseOrder{
			OrderID:  orderID,
			Account:  account,
			Stock:    stock,
			Quantity: quantity,
			Price:    price,
			status:   OrderStatusPending,
		},
	}
}

func (o *SellOrder) Execute() error {
	if err := o.Account.Portfolio.RemoveStock(o.Stock, o.Quantity); err != nil {
		o.SetStatus(OrderStatusRejected)
		return err
	}

	totalProceeds := float64(o.Quantity) * o.Price
	o.Account.Deposit(totalProceeds)
	o.SetStatus(OrderStatusExecuted)
	return nil
}
