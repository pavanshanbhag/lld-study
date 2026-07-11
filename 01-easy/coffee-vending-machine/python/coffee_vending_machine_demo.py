from coffee_vending_machine import CoffeeVendingMachine
from enums import CoffeeType, Ingredient, ToppingType
from inventory import Inventory


def main() -> None:
    inventory = Inventory()
    machine = CoffeeVendingMachine(inventory)

    print("=== Initializing Vending Machine ===")
    inventory.add_stock(Ingredient.COFFEE_BEANS, 50)
    inventory.add_stock(Ingredient.WATER, 500)
    inventory.add_stock(Ingredient.MILK, 200)
    inventory.add_stock(Ingredient.SUGAR, 100)
    inventory.add_stock(Ingredient.CARAMEL_SYRUP, 50)
    inventory.print_inventory()

    print("\n--- SCENARIO 1: Buy a Latte (Success) ---")
    machine.select_coffee(CoffeeType.LATTE, [])
    machine.insert_money(200)
    machine.insert_money(50)
    machine.dispense_coffee()
    inventory.print_inventory()

    print("\n--- SCENARIO 2: Buy Espresso (Insufficient Funds & Cancel) ---")
    machine.select_coffee(CoffeeType.ESPRESSO, [])
    machine.insert_money(100)
    machine.dispense_coffee()
    machine.cancel()
    inventory.print_inventory()

    print("\n--- SCENARIO 3: Buy Cappuccino (Out of Milk) ---")
    inventory.print_inventory()
    machine.select_coffee(
        CoffeeType.CAPPUCCINO,
        [ToppingType.CARAMEL_SYRUP, ToppingType.EXTRA_SUGAR],
    )
    machine.insert_money(300)
    machine.dispense_coffee()
    inventory.print_inventory()

    print("\n--- REFILLING AND FINAL TEST ---")
    inventory.add_stock(Ingredient.MILK, 200)
    inventory.print_inventory()
    machine.select_coffee(CoffeeType.LATTE, [ToppingType.CARAMEL_SYRUP])
    machine.insert_money(250)
    machine.dispense_coffee()
    inventory.print_inventory()


if __name__ == "__main__":
    main()
