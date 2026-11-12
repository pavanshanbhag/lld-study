from coin import Coin
from vending_machine import VendingMachine


def test_vending_machine_add_item() -> None:
    vm = VendingMachine()
    item = vm.add_item("A1", "Water", 10, 2)
    assert item.name == "Water"
    vm.select_item("A1")
    vm.insert_coin(Coin.QUARTER)
    vm.dispense()
    assert vm.inventory.stock_map["A1"] == 1


def test_unavailable_item_stays_idle() -> None:
    vm = VendingMachine()
    vm.select_item("MISSING")
    assert vm.selected_item_code is None
    assert vm.get_balance() == 0
