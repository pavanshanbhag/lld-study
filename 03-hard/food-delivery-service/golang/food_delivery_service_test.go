package fooddeliveryservice

import "testing"

func TestFoodDeliveryServiceRegisterAndPlaceOrder(t *testing.T) {
	t.Parallel()

	service := NewFoodDeliveryService()
	customer := NewCustomer("C1", "Alice", "alice@example.com", "555-0100")
	restaurant := NewRestaurant("R1", "Pizza Place", "123 Main St", nil)
	service.RegisterCustomer(customer)
	service.RegisterRestaurant(restaurant)

	item := NewMenuItem("M1", "Margherita", "Classic pizza", 12.99)
	items := []*OrderItem{NewOrderItem(item, 1)}

	order, err := service.PlaceOrder("C1", "R1", items)
	if err != nil {
		t.Fatalf("PlaceOrder: %v", err)
	}
	if order == nil || order.Customer.ID != "C1" {
		t.Fatal("expected order for customer C1")
	}
}
