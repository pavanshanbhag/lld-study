package main

import (
	"fmt"

	"decorator"
)

func main() {
	var drink decorator.Beverage = decorator.NewSimpleCoffee()
	fmt.Printf("%s $%.2f\n", drink.GetDescription(), drink.Cost())

	drink = decorator.NewMilkDecorator(drink)
	fmt.Printf("%s $%.2f\n", drink.GetDescription(), drink.Cost())

	drink = decorator.NewSugarDecorator(drink)
	fmt.Printf("%s $%.2f\n", drink.GetDescription(), drink.Cost())
}
