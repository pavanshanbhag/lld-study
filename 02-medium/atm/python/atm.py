from typing import Optional
from atm_state import IdleState
from bank_service import BankService
from card import Card
from note_dispenser import NoteDispenser100, NoteDispenser50, NoteDispenser20
from cash_dispenser import CashDispenser


class ATM:
    def __init__(self):
        self._current_state = IdleState()
        self._bank_service = BankService()
        self._current_card: Optional[Card] = None
        self._transaction_counter = 0

        c1 = NoteDispenser100(10)
        c2 = NoteDispenser50(20)
        c3 = NoteDispenser20(30)
        c1.set_next_chain(c2)
        c2.set_next_chain(c3)
        self._cash_dispenser = CashDispenser(c1)

    def change_state(self, new_state):
        self._current_state = new_state

    def set_current_card(self, card: Optional[Card]):
        self._current_card = card

    def insert_card(self, card_number: str):
        self._current_state.insert_card(self, card_number)

    def enter_pin(self, pin: str):
        self._current_state.enter_pin(self, pin)

    def select_operation(self, op, *args):
        self._current_state.select_operation(self, op, *args)

    def check_balance(self):
        balance = self._bank_service.get_balance(self._current_card)
        print(f"Your current account balance is: ${balance:.2f}")

    def withdraw_cash(self, amount: int):
        if not self._cash_dispenser.can_dispense_cash(amount):
            raise RuntimeError("Insufficient cash available in the ATM.")

        self._bank_service.withdraw_money(self._current_card, amount)

        try:
            self._cash_dispenser.dispense_cash(amount)
        except Exception as e:
            self._bank_service.deposit_money(self._current_card, amount)
            raise e

    def deposit_cash(self, amount: int):
        self._bank_service.deposit_money(self._current_card, amount)

    def get_current_card(self) -> Optional[Card]:
        return self._current_card

    def get_bank_service(self) -> BankService:
        return self._bank_service
