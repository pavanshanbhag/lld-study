package carrentalsystem

type PayPalPaymentProcessor struct{}

func NewPayPalPaymentProcessor() *PayPalPaymentProcessor {
	return &PayPalPaymentProcessor{}
}

func (p *PayPalPaymentProcessor) ProcessPayment(amount float64) error {
	return nil
}
