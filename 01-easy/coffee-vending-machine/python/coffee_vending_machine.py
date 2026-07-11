
from coffee import Coffee
from coffee_factory import CoffeeFactory
from decorators import CaramelSyrupDecorator, ExtraSugarDecorator
from enums import CoffeeType, ToppingType
from inventory import Inventory
from states import ReadyState, VendingMachineState


class CoffeeVendingMachine:
    def __init__(self, inventory: Inventory | None = None) -> None:
        self._inventory = inventory or Inventory()
        self._state: VendingMachineState = ReadyState()
        self._selected_coffee: Coffee | None = None
        self._money_inserted = 0

    def select_coffee(self, coffee_type: CoffeeType, toppings: list[ToppingType]) -> None:
        coffee = CoffeeFactory.create_coffee(coffee_type)

        for topping in toppings:
            if topping == ToppingType.EXTRA_SUGAR:
                coffee = ExtraSugarDecorator(coffee)
            elif topping == ToppingType.CARAMEL_SYRUP:
                coffee = CaramelSyrupDecorator(coffee)

        self._state.select_coffee(self, coffee)

    def insert_money(self, amount: int) -> None:
        self._state.insert_money(self, amount)

    def dispense_coffee(self) -> None:
        self._state.dispense_coffee(self)

    def cancel(self) -> None:
        self._state.cancel(self)

    def set_state(self, state: VendingMachineState) -> None:
        self._state = state

    def get_state(self) -> VendingMachineState:
        return self._state

    def set_selected_coffee(self, coffee: Coffee) -> None:
        self._selected_coffee = coffee

    def get_selected_coffee(self) -> Coffee:
        return self._selected_coffee

    def set_money_inserted(self, amount: int) -> None:
        self._money_inserted = amount

    def get_money_inserted(self) -> int:
        return self._money_inserted

    def get_inventory(self) -> Inventory:
        return self._inventory

    def reset(self) -> None:
        self._selected_coffee = None
        self._money_inserted = 0
