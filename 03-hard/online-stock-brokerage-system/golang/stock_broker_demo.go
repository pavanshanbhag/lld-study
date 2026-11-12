package onlinestockbrokeragesystem

import "fmt"

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	broker := NewStockBroker()

	// Create user and account
	user := NewUser("U001", "John Doe", "john@example.com")
	broker.CreateAccount(user, 10000.0)
	account := broker.GetAccount("A001")

	// Add stocks
	apple := NewStock("AAPL", "Apple Inc.", 150.0)
	google := NewStock("GOOGL", "Alphabet Inc.", 2000.0)
	broker.AddStock(apple)
	broker.AddStock(google)

	// Place buy orders
	buyOrder1 := NewBuyOrder("O001", account, apple, 10, 150.0)
	buyOrder2 := NewBuyOrder("O002", account, google, 5, 2000.0)

	broker.PlaceOrder(buyOrder1)
	broker.PlaceOrder(buyOrder2)

	// Give some time for orders to process
	// In a real system, you'd use proper synchronization
	logger.Info("demo", "msg", fmt.Sprint("Processing buy orders..."))

	// Place sell order
	sellOrder := NewSellOrder("O003", account, apple, 5, 160.0)
	broker.PlaceOrder(sellOrder)

	logger.Info("demo", "msg", fmt.Sprint("Processing sell order..."))

	// Print account details
	logger.Info("demo", "msg", fmt.Sprintf("Account Balance: $%.2f\n", account.GetBalance()))
	logger.Info("demo", "msg", fmt.Sprint("Portfolio Holdings:"))
	for symbol, quantity := range account.Portfolio.GetHoldings() {
		logger.Info("demo", "msg", fmt.Sprintf("%s: %d shares\n", symbol, quantity))
	}
}
