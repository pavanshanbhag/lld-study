package vending_machine

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	vm := NewVendingMachine()

	coke := NewProduct("Coke", 1.5)
	pepsi := NewProduct("Pepsi", 1.5)
	water := NewProduct("Water", 1.0)

	vm.inventory.AddProduct(coke, 5)
	vm.inventory.AddProduct(pepsi, 3)
	vm.inventory.AddProduct(water, 2)

	logger.Info("starting vending machine demo")

	logger.Info("selecting product", "name", coke.Name())
	if err := vm.SelectProduct(coke); err != nil {
		logger.Error("select failed", "err", err)
	}

	logger.Info("inserting coins")
	for range 4 {
		if err := vm.InsertCoin(QUARTER); err != nil {
			logger.Error("insert coin failed", "err", err)
		}
	}

	logger.Info("dispensing product")
	if err := vm.DispenseProduct(); err != nil {
		logger.Error("dispense failed", "err", err)
	}

	logger.Info("returning change")
	if err := vm.ReturnChange(); err != nil {
		logger.Error("return change failed", "err", err)
	}

	logger.Info("selecting product with insufficient funds", "name", pepsi.Name())
	_ = vm.SelectProduct(pepsi)
	_ = vm.InsertCoin(QUARTER)

	if err := vm.DispenseProduct(); err != nil {
		logger.Info("dispense blocked as expected", "err", err)
	}

	logger.Info("adding more coins")
	for range 4 {
		_ = vm.InsertCoin(QUARTER)
	}

	if err := vm.DispenseProduct(); err != nil {
		logger.Error("dispense failed", "err", err)
	}
	if err := vm.ReturnChange(); err != nil {
		logger.Error("return change failed", "err", err)
	}
}
