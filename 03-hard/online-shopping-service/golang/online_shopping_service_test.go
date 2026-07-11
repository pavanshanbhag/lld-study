package onlineshopping

import "testing"

func TestOnlineShoppingServiceSearchProducts(t *testing.T) {
	t.Parallel()

	service := NewOnlineShoppingService()
	product := NewProduct("P1", "Smartphone", "High-end phone", 999.99, 10)
	service.AddProduct(product)

	results := service.SearchProducts("phone")
	if len(results) != 1 || results[0].ID != "P1" {
		t.Fatalf("SearchProducts() = %+v, want one phone match", results)
	}
}

func TestOnlineShoppingServicePlaceOrder(t *testing.T) {
	t.Parallel()

	service := NewOnlineShoppingService()
	user := NewUser("U1", "Alice", "alice@example.com", "secret")
	product := NewProduct("P1", "Laptop", "Gaming laptop", 1500, 5)
	service.RegisterUser(user)
	service.AddProduct(product)

	cart := NewShoppingCart()
	cart.AddItem(product, 1)

	order, err := service.PlaceOrder(user, cart, &CreditCardPayment{})
	if err != nil {
		t.Fatalf("PlaceOrder: %v", err)
	}
	if order == nil || order.TotalAmount != 1500 {
		t.Fatalf("order total = %v, want 1500", order.TotalAmount)
	}
}
