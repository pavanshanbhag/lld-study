from atm import ATM


def test_atm_constructor() -> None:
    atm = ATM()
    assert atm.get_bank_service() is not None
    assert atm.get_current_card() is None


def test_atm_insert_card_sets_current_card() -> None:
    atm = ATM()
    atm.insert_card("1234-5678-9012-3456")
    card = atm.get_current_card()
    assert card is not None
    assert card.get_card_number() == "1234-5678-9012-3456"
