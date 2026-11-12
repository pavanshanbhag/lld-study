from coffee_vending_machine import CoffeeVendingMachine
from enums import Ingredient
from inventory import Inventory


def test_coffee_machine_with_inventory() -> None:
    inventory = Inventory()
    inventory.add_stock(Ingredient.COFFEE_BEANS, 10)
    inventory.add_stock(Ingredient.WATER, 10)
    machine = CoffeeVendingMachine(inventory)
    assert machine.get_inventory() is inventory
