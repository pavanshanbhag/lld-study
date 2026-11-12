from __future__ import annotations

import logging
from abc import ABC, abstractmethod
from typing import TYPE_CHECKING

from coin import Coin

if TYPE_CHECKING:
    from vending_machine import VendingMachine

logger = logging.getLogger(__name__)


class VendingMachineState(ABC):
    def __init__(self, machine: VendingMachine):
        self.machine = machine

    @abstractmethod
    def insert_coin(self, coin: Coin) -> None: ...

    @abstractmethod
    def select_item(self, code: str) -> None: ...

    @abstractmethod
    def dispense(self) -> None: ...

    @abstractmethod
    def refund(self) -> None: ...


class IdleState(VendingMachineState):
    def insert_coin(self, coin: Coin) -> None:
        logger.info("please select an item before inserting money")

    def select_item(self, code: str) -> None:
        if not self.machine.get_inventory().is_available(code):
            logger.warning("item not available: %s", code)
            return

        self.machine.set_selected_item_code(code)
        from states import ItemSelectedState

        self.machine.set_state(ItemSelectedState(self.machine))
        logger.info("item selected: %s", code)

    def dispense(self) -> None:
        logger.info("no item selected")

    def refund(self) -> None:
        logger.info("no money to refund")


class ItemSelectedState(VendingMachineState):
    def insert_coin(self, coin: Coin) -> None:
        self.machine.add_balance(int(coin))
        logger.info("coin inserted: %s¢ (%s)", int(coin), coin.name)

        selected_item = self.machine.get_selected_item()
        if selected_item and self.machine.get_balance() >= selected_item.price:
            logger.info("sufficient money received")
            from states import HasMoneyState

            self.machine.set_state(HasMoneyState(self.machine))

    def select_item(self, code: str) -> None:
        logger.info("item already selected")

    def dispense(self) -> None:
        logger.info("please insert sufficient money")

    def refund(self) -> None:
        self.machine.refund_balance()
        self.machine.reset()
        from states import IdleState

        self.machine.set_state(IdleState(self.machine))


class HasMoneyState(VendingMachineState):
    def insert_coin(self, coin: Coin) -> None:
        self.machine.add_balance(int(coin))
        logger.info("additional coin inserted: %s¢", int(coin))

    def select_item(self, code: str) -> None:
        logger.info("item already selected")

    def dispense(self) -> None:
        from states import DispensingState

        self.machine.set_state(DispensingState(self.machine))
        self.machine.dispense_item()

    def refund(self) -> None:
        self.machine.refund_balance()
        self.machine.reset()
        from states import IdleState

        self.machine.set_state(IdleState(self.machine))


class DispensingState(VendingMachineState):
    def insert_coin(self, coin: Coin) -> None:
        logger.info("currently dispensing")

    def select_item(self, code: str) -> None:
        logger.info("currently dispensing")

    def dispense(self) -> None:
        logger.info("dispensing in progress")

    def refund(self) -> None:
        logger.info("refund not allowed while dispensing")
