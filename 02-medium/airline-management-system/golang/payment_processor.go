package airlinemanagementsystem

import "sync"

type PaymentProcessor struct {
	mu sync.Mutex
}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

func (p *PaymentProcessor) ProcessPayment(payment *Payment) {
	p.mu.Lock()
	defer p.mu.Unlock()
	payment.ProcessPayment()
}
