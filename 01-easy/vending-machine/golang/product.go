package vending_machine

type Product struct {
	name  string
	price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{name: name, price: price}
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}
