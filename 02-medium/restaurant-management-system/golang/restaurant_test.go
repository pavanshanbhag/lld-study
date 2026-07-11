package restaurantmanagementsystem

import "testing"

func TestRestaurantPlaceOrder(t *testing.T) {
	t.Parallel()

	restaurant := NewRestaurant()
	item := NewMenuItem(1, "Burger", "Tasty burger", 9.99, true)
	restaurant.AddMenuItem(item)

	order := NewOrder(1, []MenuItem{*item}, 9.99, OrderPending)
	restaurant.PlaceOrder(order)

	restaurant.UpdateOrderStatus(1, OrderPreparing)
	if order.Status != OrderPreparing {
		t.Fatalf("order status = %v, want preparing", order.Status)
	}
}

func TestRestaurantGetMenu(t *testing.T) {
	t.Parallel()

	restaurant := NewRestaurant()
	restaurant.AddMenuItem(NewMenuItem(1, "Pizza", "Cheesy", 12.99, true))

	menu := restaurant.GetMenu()
	if len(menu) != 1 || menu[0].Name != "Pizza" {
		t.Fatalf("GetMenu() = %+v, want one pizza item", menu)
	}
}
