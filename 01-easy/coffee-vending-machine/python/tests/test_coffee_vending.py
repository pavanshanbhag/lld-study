from coffee_vending_machine import CoffeeVendingMachine
from enums import CoffeeType, Ingredient
from inventory import Inventory


def test_coffee_machine_with_inventory() -> None:
    inventory = Inventory()
    inventory.add_stock(Ingredient.COFFEE_BEANS, 10)
    inventory.add_stock(Ingredient.WATER, 10)
    machine = CoffeeVendingMachine(inventory)
    assert machine.get_inventory() is inventory


def test_dispense_espresso_deducts_stock() -> None:
    inventory = Inventory()
    inventory.add_stock(Ingredient.COFFEE_BEANS, 20)
    inventory.add_stock(Ingredient.WATER, 50)
    machine = CoffeeVendingMachine(inventory)
    machine.select_coffee(CoffeeType.ESPRESSO, [])
    machine.insert_money(150)
    machine.dispense_coffee()
    assert inventory.stock()[Ingredient.COFFEE_BEANS] == 13
    assert inventory.stock()[Ingredient.WATER] == 20
