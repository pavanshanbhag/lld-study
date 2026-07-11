package main

import (
	"fmt"

	"design-patterns/golang/adapter"
)

func main() {
	fmt.Println("Using In-House Payment Processor:")
	inHouseProcessor := &adapter.InHousePaymentProcessor{}
	checkoutService := adapter.NewCheckoutService(inHouseProcessor)
	checkoutService.ProcessCheckout(100.0, "USD")

	fmt.Println("\nUsing Legacy Gateway (via Adapter):")
	legacyGateway := &adapter.LegacyGateway{}
	legacyAdapter := adapter.NewLegacyGatewayAdapter(legacyGateway)
	checkoutService = adapter.NewCheckoutService(legacyAdapter)
	checkoutService.ProcessCheckout(150.0, "EUR")
}
