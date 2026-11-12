package onlineshopping

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	shoppingService := NewOnlineShoppingService()

	user1 := NewUser("U001", "John Doe", "john@example.com", "password123")
	user2 := NewUser("U002", "Jane Smith", "jane@example.com", "password456")
	shoppingService.RegisterUser(user1)
	shoppingService.RegisterUser(user2)

	product1 := NewProduct("P001", "Smartphone", "High-end smartphone", 999.99, 10)
	product2 := NewProduct("P002", "Laptop", "Powerful gaming laptop", 1999.99, 5)
	shoppingService.AddProduct(product1)
	shoppingService.AddProduct(product2)

	cart1 := NewShoppingCart()
	cart1.AddItem(product1, 2)
	cart1.AddItem(product2, 1)
	payment1 := &CreditCardPayment{}
	order1, err := shoppingService.PlaceOrder(user1, cart1, payment1)
	if err != nil {
		logger.Error("place order failed", "err", err)
	} else {
		logger.Info("order placed", "id", order1.ID)
	}

	searchResults := shoppingService.SearchProducts("laptop")
	logger.Info("search results", "count", len(searchResults))
	for _, product := range searchResults {
		logger.Info("product", "name", product.Name)
	}

	cart2 := NewShoppingCart()
	cart2.AddItem(searchResults[0], 1)
	payment2 := &CreditCardPayment{}
	order2, err := shoppingService.PlaceOrder(user2, cart2, payment2)
	if err != nil {
		logger.Error("place order failed", "err", err)
	} else {
		logger.Info("order placed", "id", order2.ID)
	}

	logger.Info("user order history", "user", user1.Name)
	for _, order := range user1.Orders {
		logger.Info("order", "id", order.ID, "total", order.TotalAmount, "status", order.Status)
	}
}
