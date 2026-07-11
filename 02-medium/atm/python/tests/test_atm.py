from atm import ATM


def test_atm_constructor() -> None:
    atm = ATM()
    assert atm.get_bank_service() is not None
    assert atm.get_current_card() is None
