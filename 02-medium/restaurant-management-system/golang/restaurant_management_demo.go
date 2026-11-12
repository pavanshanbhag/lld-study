package restaurantmanagementsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	restaurant := NewRestaurant()

	restaurant.AddMenuItem(NewMenuItem(1, "Burger", "Delicious burger", 9.99, true))
	restaurant.AddMenuItem(NewMenuItem(2, "Pizza", "Cheesy pizza", 12.99, true))
	restaurant.AddMenuItem(NewMenuItem(3, "Salad", "Fresh salad", 7.99, true))

	order := NewOrder(1, []MenuItem{
		*NewMenuItem(1, "Burger", "Delicious burger", 9.99, true),
		*NewMenuItem(3, "Salad", "Fresh salad", 7.99, true),
	}, 17.98, OrderPending)
	restaurant.PlaceOrder(order)

	restaurant.MakeReservation(NewReservation(1, "John Doe", "1234567890", 4, time.Now()))
	restaurant.ProcessPayment(NewPayment(1, 17.98, CreditCard, PaymentPending))

	restaurant.UpdateOrderStatus(1, OrderPreparing)
	restaurant.UpdateOrderStatus(1, OrderReady)
	restaurant.UpdateOrderStatus(1, OrderCompleted)

	restaurant.AddStaff(NewStaff(1, "Alice", "Manager", "9876543210"))
	restaurant.AddStaff(NewStaff(2, "Bob", "Chef", "5432109876"))

	for _, item := range restaurant.GetMenu() {
		logger.Info("menu item", "name", item.Name, "price", item.Price)
	}
}
