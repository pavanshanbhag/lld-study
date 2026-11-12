package coffeevendingmachine

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	coffeeMachine := NewCoffeeMachine()

	for _, item := range coffeeMachine.Menu() {
		logger.Info("menu item", "item", item)
	}

	coffeeTypes := []string{"Espresso", "Cappuccino", "Latte"}
	payments := []float64{3.0, 3.5, 4.0}

	for i, coffeeType := range coffeeTypes {
		selectedCoffee, err := coffeeMachine.SelectCoffee(coffeeType)
		if err != nil {
			logger.Error("select failed", "coffee", coffeeType, "err", err)
			continue
		}
		payment := NewPayment(payments[i])
		logger.Info("ordering", "coffee", coffeeType)
		change, err := coffeeMachine.DispenseCoffee(selectedCoffee, payment)
		if err != nil {
			logger.Error("dispense failed", "err", err)
			continue
		}
		logger.Info("dispensed", "coffee", coffeeType, "change", change)
	}
}
