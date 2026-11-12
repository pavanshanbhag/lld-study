package vending_machine

import "fmt"

type Inventory struct {
	products map[*Product]int
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[*Product]int)}
}

func (inv *Inventory) AddProduct(product *Product, quantity int) {
	inv.products[product] = quantity
}

func (inv *Inventory) IsAvailable(product *Product) bool {
	qty, exists := inv.products[product]
	return exists && qty > 0
}

func (inv *Inventory) Quantity(product *Product) int {
	return inv.products[product]
}

func (inv *Inventory) Dispense(product *Product) error {
	if !inv.IsAvailable(product) {
		return fmt.Errorf("product not available: %s", product.name)
	}
	inv.products[product]--
	return nil
}
