from __future__ import annotations

import logging

from coin import Coin
from inventory import Inventory
from item import Item
from states import IdleState, VendingMachineState

logger = logging.getLogger(__name__)


class VendingMachine:
    def __init__(self) -> None:
        self.inventory = Inventory()
        self.current_state: VendingMachineState = IdleState(self)
        self.balance = 0
        self.selected_item_code: str | None = None

    def insert_coin(self, coin: Coin) -> None:
        self.current_state.insert_coin(coin)

    def add_item(self, code: str, name: str, price: int, quantity: int) -> Item:
        item = Item(code, name, price)
        self.inventory.add_item(code, item, quantity)
        return item

    def select_item(self, code: str) -> None:
        self.current_state.select_item(code)

    def dispense(self) -> None:
        self.current_state.dispense()

    def dispense_item(self) -> None:
        item = self.inventory.get_item(self.selected_item_code)
        if item is None:
            logger.warning("no selected item to dispense")
            self.reset()
            self.set_state(IdleState(self))
            return
        if self.balance >= item.price:
            self.inventory.reduce_stock(self.selected_item_code)
            self.balance -= item.price
            logger.info("dispensed %s", item.name)
            if self.balance > 0:
                logger.info("returning change: %s", self.balance)
        self.reset()
        self.set_state(IdleState(self))

    def refund_balance(self) -> None:
        logger.info("refunding: %s", self.balance)
        self.balance = 0

    def reset(self) -> None:
        self.selected_item_code = None
        self.balance = 0

    def add_balance(self, value: int) -> None:
        self.balance += value

    def get_selected_item(self) -> Item | None:
        return self.inventory.get_item(self.selected_item_code)

    def set_selected_item_code(self, code: str) -> None:
        self.selected_item_code = code

    def set_state(self, state: VendingMachineState) -> None:
        self.current_state = state

    def get_inventory(self) -> Inventory:
        return self.inventory

    def get_balance(self) -> int:
        return self.balance
