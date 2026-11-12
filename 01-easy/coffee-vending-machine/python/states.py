from __future__ import annotations

import logging
from abc import ABC, abstractmethod
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from coffee import Coffee
    from coffee_vending_machine import CoffeeVendingMachine

logger = logging.getLogger(__name__)


class VendingMachineState(ABC):
    @abstractmethod
    def select_coffee(self, machine: CoffeeVendingMachine, coffee: Coffee): ...

    @abstractmethod
    def insert_money(self, machine: CoffeeVendingMachine, amount: int): ...

    @abstractmethod
    def dispense_coffee(self, machine: CoffeeVendingMachine): ...

    @abstractmethod
    def cancel(self, machine: CoffeeVendingMachine): ...


class ReadyState(VendingMachineState):
    def select_coffee(self, machine: CoffeeVendingMachine, coffee: Coffee):
        machine.set_selected_coffee(coffee)
        machine.set_state(SelectingState())
        logger.info("%s selected. Price: %s", coffee.get_coffee_type(), coffee.get_price())

    def insert_money(self, machine: CoffeeVendingMachine, amount: int):
        logger.info("Please select a coffee first.")

    def dispense_coffee(self, machine: CoffeeVendingMachine):
        logger.info("Please select and pay first.")

    def cancel(self, machine: CoffeeVendingMachine):
        logger.info("Nothing to cancel.")


class SelectingState(VendingMachineState):
    def select_coffee(self, machine: CoffeeVendingMachine, coffee: Coffee):
        logger.info("Already selected. Please pay or cancel.")

    def insert_money(self, machine: CoffeeVendingMachine, amount: int):
        machine.set_money_inserted(machine.get_money_inserted() + amount)
        logger.info("Inserted %s. Total: %s", amount, machine.get_money_inserted())
        if machine.get_money_inserted() >= machine.get_selected_coffee().get_price():
            machine.set_state(PaidState())

    def dispense_coffee(self, machine: CoffeeVendingMachine):
        logger.info("Please insert enough money first.")

    def cancel(self, machine: CoffeeVendingMachine):
        logger.info("Transaction cancelled. Refunding %s", machine.get_money_inserted())
        machine.reset()
        machine.set_state(ReadyState())


class PaidState(VendingMachineState):
    def select_coffee(self, machine: CoffeeVendingMachine, coffee: Coffee):
        logger.info("Already paid. Please dispense or cancel.")

    def insert_money(self, machine: CoffeeVendingMachine, amount: int):
        machine.set_money_inserted(machine.get_money_inserted() + amount)
        logger.info("Additional %s inserted. Total: %s", amount, machine.get_money_inserted())

    def dispense_coffee(self, machine: CoffeeVendingMachine):
        inventory = machine.get_inventory()
        coffee = machine.get_selected_coffee()

        if not inventory.has_ingredients(coffee.get_recipe()):
            logger.info("Sorry, we are out of ingredients. Refunding your money.")
            logger.info("Refunding %s", machine.get_money_inserted())
            machine.reset()
            machine.set_state(OutOfIngredientState())
            return

        inventory.deduct_ingredients(coffee.get_recipe())
        coffee.prepare()

        change = machine.get_money_inserted() - coffee.get_price()
        if change > 0:
            logger.info("Here's your change: %s", change)

        machine.reset()
        machine.set_state(ReadyState())

    def cancel(self, machine: CoffeeVendingMachine):
        logger.info("Transaction cancelled. Refunding %s", machine.get_money_inserted())
        machine.reset()
        machine.set_state(ReadyState())


class OutOfIngredientState(VendingMachineState):
    def select_coffee(self, machine: CoffeeVendingMachine, coffee: Coffee):
        logger.info("Sorry, we are sold out.")

    def insert_money(self, machine: CoffeeVendingMachine, amount: int):
        logger.info("Sorry, we are sold out. Money refunded.")

    def dispense_coffee(self, machine: CoffeeVendingMachine):
        logger.info("Sorry, we are sold out.")

    def cancel(self, machine: CoffeeVendingMachine):
        logger.info("Refunding %s", machine.get_money_inserted())
        machine.reset()
        machine.set_state(ReadyState())
