package coffeevendingmachine

import (
	"fmt"
	"sync"
)

type CoffeeMachine struct {
	coffeeMenu  []*Coffee
	ingredients map[string]*Ingredient
	mu          sync.Mutex
}

func NewCoffeeMachine() *CoffeeMachine {
	cm := &CoffeeMachine{
		coffeeMenu:  make([]*Coffee, 0),
		ingredients: make(map[string]*Ingredient),
	}
	cm.initializeIngredients()
	cm.initializeCoffeeMenu()
	return cm
}

func (cm *CoffeeMachine) initializeIngredients() {
	cm.ingredients["Coffee"] = NewIngredient("Coffee", 10)
	cm.ingredients["Water"] = NewIngredient("Water", 10)
	cm.ingredients["Milk"] = NewIngredient("Milk", 10)
}

func (cm *CoffeeMachine) initializeCoffeeMenu() {
	espressoRecipe := make(map[*Ingredient]int)
	espressoRecipe[cm.ingredients["Coffee"]] = 1
	espressoRecipe[cm.ingredients["Water"]] = 1
	cm.coffeeMenu = append(cm.coffeeMenu, NewCoffee("Espresso", 2.5, espressoRecipe))

	cappuccinoRecipe := make(map[*Ingredient]int)
	cappuccinoRecipe[cm.ingredients["Coffee"]] = 1
	cappuccinoRecipe[cm.ingredients["Water"]] = 1
	cappuccinoRecipe[cm.ingredients["Milk"]] = 1
	cm.coffeeMenu = append(cm.coffeeMenu, NewCoffee("Cappuccino", 3.5, cappuccinoRecipe))

	latteRecipe := make(map[*Ingredient]int)
	latteRecipe[cm.ingredients["Coffee"]] = 1
	latteRecipe[cm.ingredients["Water"]] = 1
	latteRecipe[cm.ingredients["Milk"]] = 2
	cm.coffeeMenu = append(cm.coffeeMenu, NewCoffee("Latte", 4.0, latteRecipe))
}

func (cm *CoffeeMachine) Menu() []string {
	items := make([]string, 0, len(cm.coffeeMenu))
	for _, coffee := range cm.coffeeMenu {
		items = append(items, fmt.Sprintf("%s - $%.2f", coffee.Name(), coffee.Price()))
	}
	return items
}

func (cm *CoffeeMachine) SelectCoffee(coffeeName string) (*Coffee, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, coffee := range cm.coffeeMenu {
		if coffee.Name() == coffeeName {
			return coffee, nil
		}
	}
	return nil, fmt.Errorf("unknown coffee: %s", coffeeName)
}

func (cm *CoffeeMachine) DispenseCoffee(coffee *Coffee, payment *Payment) (float64, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if coffee == nil {
		return 0, fmt.Errorf("invalid coffee selection")
	}

	if payment.amount < coffee.Price() {
		return 0, fmt.Errorf("insufficient payment for %s", coffee.Name())
	}

	if !cm.hasEnoughIngredients(coffee) {
		return 0, fmt.Errorf("insufficient ingredients to make %s", coffee.Name())
	}

	cm.updateIngredients(coffee)
	return payment.amount - coffee.Price(), nil
}

func (cm *CoffeeMachine) hasEnoughIngredients(coffee *Coffee) bool {
	for ingredient, requiredQuantity := range coffee.Recipe() {
		if ingredient.Quantity() < requiredQuantity {
			return false
		}
	}
	return true
}

func (cm *CoffeeMachine) updateIngredients(coffee *Coffee) {
	for ingredient, requiredQuantity := range coffee.Recipe() {
		ingredient.UpdateQuantity(-requiredQuantity)
	}
}

func (cm *CoffeeMachine) IngredientQuantity(name string) int {
	ing, ok := cm.ingredients[name]
	if !ok {
		return 0
	}
	return ing.Quantity()
}
