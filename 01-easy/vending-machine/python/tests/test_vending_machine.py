from coin import Coin
from vending_machine import VendingMachine


def test_vending_machine_add_item() -> None:
    vm = VendingMachine()
    item = vm.add_item("A1", "Water", 10, 2)
    assert item.get_name() == "Water"
    vm.select_item("A1")
    vm.insert_coin(Coin.QUARTER)
    vm.dispense()
