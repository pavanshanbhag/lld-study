from typing import TYPE_CHECKING
from item_states import AvailableState

if TYPE_CHECKING:
    from member import Member
    from library_item import LibraryItem
    from item_states import ItemState
    from transaction_service import TransactionService


class BookCopy:
    def __init__(self, copy_id: str, item: 'LibraryItem', transaction_service: 'TransactionService'):
        self.id = copy_id
        self.item = item
        self._transaction_service = transaction_service
        self.current_state: 'ItemState' = AvailableState()
        item.add_copy(self)

    def checkout(self, member: 'Member') -> None:
        self.current_state.checkout(self, member)

    def return_item(self) -> None:
        self.current_state.return_item(self)

    def place_hold(self, member: 'Member') -> None:
        self.current_state.place_hold(self, member)

    def set_state(self, state: 'ItemState') -> None:
        self.current_state = state

    def get_id(self) -> str:
        return self.id

    def get_item(self) -> 'LibraryItem':
        return self.item

    def get_transaction_service(self) -> 'TransactionService':
        return self._transaction_service

    def is_available(self) -> bool:
        return isinstance(self.current_state, AvailableState)
